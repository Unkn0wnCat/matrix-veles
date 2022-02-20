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

	// RoomID is the rooms ID
	RoomID string `yaml:"roomID" bson:"room_id"`

	// Debug specifies if the bot shall run in dry run mode
	Debug bool `yaml:"debug" bson:"debug"`

	/*
		HashCheckMode specifies the mode the bot should operate under in this room

		HashCheck-Modes:
		 0. Silent Mode (Don't do anything)
		 1. Notify Mode (Message Room Admins & Mods)
		 2. Mute Mode (Remove message, notify admin & mute user)
		 3. Ban Mode (Remove message, notify admin & ban user)
	*/
	HashCheckMode uint8 `yaml:"mode" bson:"hash_check_mode"`
}
