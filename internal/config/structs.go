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

import "go.mongodb.org/mongo-driver/bson/primitive"

// RoomConfigTree is a map from string to RoomConfig
type RoomConfigTree map[string]RoomConfig

// RoomConfig is the configuration attached to every joined room
type RoomConfig struct {
	ID primitive.ObjectID `bson:"_id"`

	// Active tells if the bot is active in this room (Set to false on leave/kick/ban)
	Active bool `yaml:"active" bson:"active"`

	// Deactivate can be set by an admin to disable the bot for a room
	Deactivate bool `yaml:"deactivate" bson:"deactivate"`

	// Name is fetched regularly from the room state
	Name string `yaml:"name" bson:"name"`

	// RoomID is the rooms ID
	RoomID string `yaml:"roomID" bson:"room_id"`

	// Debug specifies if the bot shall run in dry run mode
	Debug bool `yaml:"debug" bson:"debug"`

	// AlertChannel is currently unused
	AlertChannel *string `bson:"alert_channel"`

	// AdminPowerLevel specifies the power-level a user has to have to manage the room
	AdminPowerLevel int `bson:"admin_power_level"`

	// HashChecker contains configuration specific to the hash-checker
	HashChecker HashCheckerConfig `bson:"hash_checker"`

	Admins []string `bson:"admins"`
}

type HashCheckerConfig struct {
	// NoticeToChat specifies weather or not to post a public notice to chat
	NoticeToChat bool `bson:"chat_notice"`

	// NotificationPowerLevel is currently unused
	NotificationPowerLevel int `yaml:"notification_level" bson:"notification_level"`

	/*
		HashCheckMode specifies the mode the bot should operate under in this room

		HashCheck-Modes:
		 0. Notice Mode (Post notice)
		 1. Delete Mode (Remove message, post notice)
		 2. Mute Mode (Remove message, post notice & mute user)
		 3. Ban Mode (Remove message, post notice & ban user)
	*/
	HashCheckMode uint8 `yaml:"mode" bson:"hash_check_mode"`

	// SubscribedLists contains the lists this room is subscribed to
	SubscribedLists []*primitive.ObjectID `bson:"subscribed_lists" json:"subscribed_lists"`
}
