/*
 * Copyright © 2022 Kevin Kandlbinder.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package config

import (
	"context"
	"errors"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/Unkn0wnCat/matrix-veles/internal/tracer"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"maunium.net/go/mautrix/id"
	"sync"
)

var (
	roomConfigWg sync.WaitGroup
)

// SetRoomConfigActive updates the active state for a given room
func SetRoomConfigActive(id string, active bool) {
	// Lock room config system to prevent race conditions
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfig := GetRoomConfig(id)
	roomConfig.Active = active

	err := SaveRoomConfig(&roomConfig)
	if err != nil {
		log.Panicf("Error writing room config to database: %v", err)
	}

	// Unlock room config system
	roomConfigWg.Done()
}

// GetRoomConfig returns the RoomConfig linked to the specified ID
func GetRoomConfig(id string) RoomConfig {
	config, err := GetRoomConfigByRoomID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return AddRoomConfig(id)
		}
	}

	return *config
}

// RoomConfigInitialUpdate updates all RoomConfig entries to set activity and create blank configs
func RoomConfigInitialUpdate(ids []id.RoomID, parentCtx context.Context) {
	ctx, span := tracer.Tracer.Start(parentCtx, "initial_room_config_update")
	defer span.End()

	_, dbGetAllSpan := tracer.Tracer.Start(ctx, "db_fetch_room_configs")
	database := db.DbClient.Database(viper.GetString("bot.mongo.database"))

	cursor, err := database.Collection("rooms").Find(context.TODO(), bson.D{}, nil)
	if err != nil {
		dbGetAllSpan.RecordError(err)
		log.Panicf("Error querying room configs: %v", err)
	}

	var roomConfigs []RoomConfig

	err = cursor.All(context.TODO(), &roomConfigs)
	if err != nil {
		dbGetAllSpan.RecordError(err)
		log.Panicf("Error querying room configs: %v", err)
	}

	dbGetAllSpan.End()

	ctx2, updateConfigs := tracer.Tracer.Start(ctx, "update_room_configs")

	activeRooms := make(map[string]bool)

	// Set all active states to "false" for a blank start
	for _, roomConfig := range roomConfigs {
		activeRooms[roomConfig.RoomID] = false
	}

	_, updateRooms := tracer.Tracer.Start(ctx2, "update_joined_rooms")
	// Go over all joined rooms
	for _, roomID := range ids {
		activeRooms[roomID.String()] = true

		GetRoomConfig(roomID.String())
	}
	updateRooms.End()

	_, saveConfigs := tracer.Tracer.Start(ctx2, "save_room_configs")
	for roomID, isActive := range activeRooms {
		SetRoomConfigActive(roomID, isActive)
	}
	saveConfigs.End()

	updateConfigs.End()
}

func AddRoomConfig(id string) RoomConfig {
	// Lock room config system to prevent race conditions
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	config := GetDefaultRoomConfig()
	config.RoomID = id

	err := SaveRoomConfig(&config)
	if err != nil {
		log.Panicf("Error writing room config to database: %v", err)
	}

	// Unlock room config system
	roomConfigWg.Done()

	return config
}

func SaveRoomConfig(roomConfig *RoomConfig) error {
	database := db.DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.Replace().SetUpsert(true)

	filter := bson.D{{"room_id", roomConfig.RoomID}}

	_, err := database.Collection(viper.GetString("bot.mongo.collection.rooms")).ReplaceOne(context.TODO(), filter, roomConfig, opts)

	return err
}

func GetRoomConfigByRoomID(id string) (*RoomConfig, error) {
	database := db.DbClient.Database(viper.GetString("bot.mongo.database"))

	res := database.Collection(viper.GetString("bot.mongo.collection.rooms")).FindOne(context.TODO(), bson.D{{"room_id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	object := GetDefaultRoomConfig()

	err := res.Decode(&object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func GetDefaultRoomConfig() RoomConfig {
	return RoomConfig{
		ID:              primitive.NewObjectID(),
		Active:          true,
		RoomID:          "",
		Debug:           false,
		AdminPowerLevel: 100,
		HashChecker: HashCheckerConfig{
			NoticeToChat:           true,
			NotificationPowerLevel: 50,
			HashCheckMode:          1,
		},
	}
}
