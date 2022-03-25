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
	"github.com/Unkn0wnCat/matrix-veles/internal/config"
	"github.com/Unkn0wnCat/matrix-veles/internal/tracer"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	messageEventsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "bot",
		Name:      "processed_message_events_total",
		Help:      "The total number of processed message events",
	})
	memberEventsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "bot",
		Name:      "processed_member_events_total",
		Help:      "The total number of processed member events",
	})
	filesProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "bot",
		Name:      "processed_files_total",
		Help:      "The total number of processed files",
	})
	commandsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "bot",
		Name:      "processed_commands_total",
		Help:      "The total number of processed commands",
	})
	joinedRooms = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "bot",
		Name:      "joined_rooms",
		Help:      "Number of rooms the bot has joined",
	})
)

// Run starts the bot, blocking until an interrupt or SIGTERM is received
func Run() {
	// Set up signal channel for controlled shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Save Timestamp for filtering
	startTs := time.Now().Unix()

	checkConfig()

	log.Printf("matrix-veles has started.")

	// Initialize client, this does not check access key!
	matrixClient, err := mautrix.NewClient(
		viper.GetString("bot.homeserver_url"),
		id.NewUserID(viper.GetString("bot.username"), viper.GetString("bot.homeserver")),
		viper.GetString("bot.accessKey"))
	if err != nil {
		log.Printf("matrix-veles couldn't initialize matrix client, please check credentials")
		log.Fatal(err)
		return
	}

	// If no accessKey is set, perform login
	if viper.GetString("bot.accessKey") == "" {
		performLogin(matrixClient)
	}

	// Set up sync handlers for invites and messages
	syncer := matrixClient.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, handleMessageEvent(matrixClient, startTs))
	syncer.OnEventType(event.StateMember, handleMemberEvent(matrixClient, startTs))

	// Set up async tasks
	go startSync(matrixClient)
	go doInitialUpdate(matrixClient)
	go doAdminStateUpdate(matrixClient)

	go func() {
		ticker := time.NewTicker(5 * time.Minute)

		for {
			go doAdminStateUpdate(matrixClient)
			<-ticker.C
		}
	}()

	<-c
	log.Printf("Shutting down...")

	matrixClient.StopSync()

	tracer.TraceProvider.Shutdown(tracer.Ctx)

	log.Printf("Goodbye!")

	tracer.Cancel()
	os.Exit(0)
}

// checkConfig applies constraints to the configuration and exits the program on violation
func checkConfig() {
	// Both homeserver and username are required!
	if viper.GetString("bot.homeserver") == "" || viper.GetString("bot.username") == "" {
		log.Printf("matrix-veles is missing user identification (homeserver / username)")
		os.Exit(1)
		return
	}

	// Either accessKey or password are required
	if viper.GetString("bot.accessKey") == "" && viper.GetString("bot.password") == "" {
		log.Printf("matrix-veles is missing user credentials (access-key / password)")
		log.Printf("Please provide either an access-key or password")
		os.Exit(1)
		return
	}
}

// performLogin logs in the mautrix.Client using the username and password from config
func performLogin(matrixClient *mautrix.Client) {
	res, err := matrixClient.Login(&mautrix.ReqLogin{
		Type:                     "m.login.password",
		Identifier:               mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: viper.GetString("bot.username")},
		Password:                 viper.GetString("bot.password"),
		StoreCredentials:         true,
		InitialDeviceDisplayName: "github.com/Unkn0wnCat/matrix-veles",
	})
	if err != nil {
		log.Printf("matrix-veles couldn't sign in, please check credentials")
		log.Fatal(err)
		return
	}

	accessToken := res.AccessToken

	// Save accessKey to configuration
	viper.Set("bot.accessKey", accessToken)
	err = viper.WriteConfig()
	if err != nil {
		log.Printf("matrix-veles could not save the accessKey to config")
		log.Fatal(err)
		return
	}
}

// doInitialUpdate updates the config right after startup to catch up with joined/left rooms
func doInitialUpdate(matrixClient *mautrix.Client) {
	ctx, span := tracer.Tracer.Start(tracer.Ctx, "initial_update")
	defer span.End()

	_, requestSpan := tracer.Tracer.Start(ctx, "request_joined_rooms")
	resp, err := matrixClient.JoinedRooms()
	if err != nil {
		log.Printf("matrix-veles could not read joined rooms, something is horribly wrong")
		log.Fatalln(err)
	}
	requestSpan.End()

	joinedRooms.Set(float64(len(resp.JoinedRooms)))

	// Hand-off list to config helper
	config.RoomConfigInitialUpdate(resp.JoinedRooms, ctx)
}

func doAdminStateUpdate(matrixClient *mautrix.Client) {
	ctx, span := tracer.Tracer.Start(tracer.Ctx, "admin_state_update")
	defer span.End()

	_, requestSpan := tracer.Tracer.Start(ctx, "request_joined_rooms")
	resp, err := matrixClient.JoinedRooms()
	if err != nil {
		log.Printf("matrix-veles could not read joined rooms, something is horribly wrong")
		log.Fatalln(err)
	}
	requestSpan.End()

	for _, roomId := range resp.JoinedRooms {
		_, processSpan := tracer.Tracer.Start(ctx, "process_room")
		processSpan.SetAttributes(attribute.String("room_id", roomId.String()))

		powerLevels, err := GetRoomPowerLevelState(matrixClient, roomId)
		if err != nil {
			log.Printf("Failed to get power levels for %s - %v", roomId.String(), err)
			processSpan.RecordError(err)
			processSpan.End()
			continue
		}

		roomConfig := config.GetRoomConfig(roomId.String())

		var admins []string

		for user, powerLevel := range powerLevels.Users {
			if powerLevel >= roomConfig.AdminPowerLevel {
				admins = append(admins, user)
			}
		}

		roomConfig = config.GetRoomConfig(roomId.String())
		roomConfig.Admins = admins
		err = config.SaveRoomConfig(&roomConfig)
		if err != nil {
			processSpan.RecordError(err)
			processSpan.End()
			continue
		}

		processSpan.End()
	}
}

// handleMessageEvent wraps message handler taking the mautrix.Client and start timestamp as parameters
func handleMessageEvent(matrixClient *mautrix.Client, startTs int64) mautrix.EventHandler {
	return func(source mautrix.EventSource, evt *event.Event) {
		ctx, span := tracer.Tracer.Start(tracer.Ctx, "handle_message_event")
		defer span.End()

		if evt.Timestamp < (startTs * 1000) {
			// Ignore old events
			return
		}

		// Cast event to correct event type
		content, ok := evt.Content.Parsed.(*event.MessageEventContent)

		defer messageEventsProcessed.Inc()

		if !ok {
			log.Println("Uh oh, could not typecast m.room.member event content...")
			return
		}

		username, _, err := matrixClient.UserID.Parse()
		if err != nil {
			log.Panicln("Invalid user id in client")
		}

		if content.URL != "" {
			span.SetAttributes(attribute.Bool("has_attachment", true))
			// This has an attachment!
			handleHashing(content, evt, matrixClient, ctx) // -> handleHashing.go
			return
		}
		span.SetAttributes(attribute.Bool("has_attachment", false))

		// No attachment, is this a command?
		if !strings.HasPrefix(content.Body, "!"+username) &&
			!strings.HasPrefix(content.Body, "@"+username) &&
			!(strings.HasPrefix(content.Body, username) && strings.HasPrefix(content.FormattedBody, "<a href=\"https://matrix.to/#/"+matrixClient.UserID.String()+"\">")) {
			span.SetAttributes(attribute.Bool("is_command", false))
			return
		}

		span.SetAttributes(attribute.Bool("is_command", true))

		// It is a command!
		handleCommand(content.Body, evt.Sender, evt.RoomID, matrixClient, ctx) // -> commandParser.go
	}
}

// handleMemberEvent wraps m.room.member (invite, join, leave, ban etc.) handler taking the mautrix.Client and start timestamp as parameters
func handleMemberEvent(matrixClient *mautrix.Client, startTs int64) func(source mautrix.EventSource, evt *event.Event) {
	return func(source mautrix.EventSource, evt *event.Event) {
		if *evt.StateKey != matrixClient.UserID.String() {
			return
		} // This does not concern us, as we are not the subject
		if evt.Timestamp < (startTs * 1000) {
			// Ignore old events, TODO: Handle missed invites.
			return
		}

		defer memberEventsProcessed.Inc()

		// Cast event to correct event type
		content, ok := evt.Content.Parsed.(*event.MemberEventContent)

		if !ok {
			log.Println("Uh oh, could not typecast m.room.member event content...")
			return
		}

		// If it is an invite, accept it
		if content.Membership == event.MembershipInvite {
			doAcceptInvite(matrixClient, evt.RoomID)
			return
		}

		// If it is our join event, set room to active
		if content.Membership == event.MembershipJoin {
			config.SetRoomConfigActive(evt.RoomID.String(), true)
			joinedRooms.Inc()
			return
		}

		// If we left or got banned, set room to inactive
		if content.Membership.IsLeaveOrBan() {
			config.SetRoomConfigActive(evt.RoomID.String(), false)
			joinedRooms.Dec()
			return
		}
	}
}

// startSync starts the mautrix.Client sync for receiving events
func startSync(matrixClient *mautrix.Client) {
	err := matrixClient.Sync()
	if err != nil {
		log.Printf("matrix-veles has encountered a fatal error whilst syncing")
		log.Println(err)
		os.Exit(2)
	}
	log.Println("sync exited.")
}

// formattedMessage is a helper struct for sending HTML content
type formattedMessage struct {
	Type          string `json:"msgtype"`
	Body          string `json:"body"`
	Format        string `json:"format"`
	FormattedBody string `json:"formatted_body"`
}
