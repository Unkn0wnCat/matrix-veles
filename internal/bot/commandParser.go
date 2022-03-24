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

package bot

import (
	"context"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/Unkn0wnCat/matrix-veles/internal/tracer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel/attribute"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
	"strings"
)

// handleCommand takes a command, parses it and executes any actions it implies
func handleCommand(command string, sender id.UserID, id id.RoomID, client *mautrix.Client, parentCtx context.Context) {
	ctx, span := tracer.Tracer.Start(parentCtx, "handle_command")
	defer span.End()

	_, prepSpan := tracer.Tracer.Start(ctx, "command_prepare")

	myUsername, _, err := client.UserID.Parse()
	if err != nil {
		prepSpan.RecordError(err)
		log.Panicln("Invalid user id in client")
	}

	defer commandsProcessed.Inc()

	command = strings.TrimPrefix(command, "!")        // Remove !
	command = strings.TrimPrefix(command, "@")        // Remove @
	command = strings.TrimPrefix(command, myUsername) // Remove our own username
	command = strings.TrimPrefix(command, ":")        // Remove : (as in "@soccerbot:")
	command = strings.TrimSpace(command)

	prepSpan.End()

	span.SetAttributes(attribute.String("veles.command_handler.command", command))

	// TODO: Remove this, it is debug!
	log.Println(command)

	// Is this a help command?
	if strings.HasPrefix(command, "help") {
		commandHelp(sender, id, client)
		return
	}

	if strings.HasPrefix(command, "link ") {
		commandLink(sender, id, client, strings.TrimPrefix(command, "link "))
		return
	}

	// No match :( - display help
	commandHelp(sender, id, client)
	return
}

func commandLink(userId id.UserID, id id.RoomID, client *mautrix.Client, linkId string) {
	linkId = strings.Trim(linkId, " \n\r")

	linkIdP, err := primitive.ObjectIDFromHex(linkId)
	if err != nil {
		_, _ = client.SendNotice(id, "Invalid Link ID")
		return
	}

	user, err := db.GetUserByID(linkIdP)
	if err != nil {
		_, _ = client.SendNotice(id, "Invalid Link ID")
		return
	}

	err = user.ValidateMXID(userId.String())
	if err != nil {
		_, _ = client.SendNotice(id, "No matching link request pending")
		return
	}

	err = db.SaveUser(user)
	if err != nil {
		_, _ = client.SendNotice(id, "Database error")
		return
	}

	_, _ = client.SendNotice(id, "Successfully linked")
}

func commandHelp(_ id.UserID, id id.RoomID, client *mautrix.Client) {
	// TODO: Improve help message

	// Ignore errors as we can't do anything about them, the user will probably retry
	_, _ = client.SendNotice(id, "matrix-veles help\n\n!veles-bot help - shows this help")
}
