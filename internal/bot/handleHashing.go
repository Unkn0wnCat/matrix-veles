package bot

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Unkn0wnCat/matrix-veles/internal/config"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"github.com/Unkn0wnCat/matrix-veles/internal/tracer"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/attribute"
	"io"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
)

var (
	filesFlagged = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "bot",
		Subsystem: "hash_handler",
		Name:      "flagged_files_total",
		Help:      "The total number of files found to be matching a hash",
	})
	filesCleared = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "bot",
		Subsystem: "hash_handler",
		Name:      "cleared_files_total",
		Help:      "The total number of files found to not be matching a hash",
	})
)

// handleHashing hashes and checks a message, taking configured actions on match
func handleHashing(content *event.MessageEventContent, evt *event.Event, matrixClient *mautrix.Client, parentCtx context.Context) {
	ctx, span := tracer.Tracer.Start(parentCtx, "handle_hashing")
	defer span.End()

	url, err := content.URL.Parse()
	if err != nil {
		span.RecordError(err)
		log.Printf("Error: Could not parse Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	_, dlSpan := tracer.Tracer.Start(ctx, "download_and_hash_attachment")

	reader, err := matrixClient.Download(url)
	if err != nil {
		log.Printf("Error: Could not read file from Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	dlSpan.AddEvent("reader opened")

	defer func(reader io.ReadCloser) { _ = reader.Close() }(reader)

	hashWriter := sha512.New()

	dlSpan.AddEvent("hash writer ready")

	if _, err = io.Copy(hashWriter, reader); err != nil {
		dlSpan.RecordError(err)
		log.Printf("Error: Could not hash file from Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	dlSpan.End()

	sum := hex.EncodeToString(hashWriter.Sum(nil))

	_, rcSpan := tracer.Tracer.Start(ctx, "db_fetch_room_config")

	// Fetch room configuration for adjusting behaviour
	roomConfig := config.GetRoomConfig(evt.RoomID.String())

	rcSpan.End()

	defer filesProcessed.Inc()

	_, hashCheckSpan := tracer.Tracer.Start(ctx, "check_hash")

	hashObj, err := db.GetEntryByHash(sum)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			filesCleared.Inc()
			hashCheckSpan.AddEvent("hash not in database - file cleared")
			hashCheckSpan.SetAttributes(attribute.Bool("veles.hash_checker.hash_match", false))
			hashCheckSpan.SetAttributes(attribute.Bool("veles.hash_checker.illegal_content", false))
			hashCheckSpan.End()
			if roomConfig.Debug {
				matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG - This file is not on the hashlist: %s", sum))
			}
			return
		}
		if roomConfig.Debug {
			matrixClient.SendNotice(evt.RoomID, "DEBUG - Failed to check file. See log.")
		}
		hashCheckSpan.RecordError(err)
		fmt.Printf("Error trying to check database: %v", err)
		return
	}

	hashCheckSpan.AddEvent("hash found in database")
	hashCheckSpan.SetAttributes(attribute.Bool("veles.hash_checker.hash_match", true))

	if roomConfig.Debug {
		matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG !!! This file is on the hashlist: %s", sum))
		jsonVal, _ := json.Marshal(hashObj)
		matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG:\n%s", makeFancyJSON(jsonVal)))
	}

	if !checkSubscription(&roomConfig, hashObj) {
		hashCheckSpan.AddEvent("room not subscribed to hash list - ignoring")
		hashCheckSpan.SetAttributes(attribute.Bool("veles.hash_checker.illegal_content", false))
		hashCheckSpan.End()
		return
	}

	hashCheckSpan.SetAttributes(attribute.Bool("veles.hash_checker.illegal_content", true))
	hashCheckSpan.End()

	log.Printf("Illegal content detected in room %s!", roomConfig.RoomID)

	handleIllegalContent(evt, matrixClient, hashObj, roomConfig, ctx)
}

// makeFancyJSON formats / indents a JSON string
func makeFancyJSON(input []byte) string {
	var buf bytes.Buffer
	json.Indent(&buf, input, "", "  ")
	return buf.String()
}

// checkSubscription checks if the room is subscribed to one of hashObjs lists
func checkSubscription(roomConfig *config.RoomConfig, hashObj *model.DBEntry) bool {
	if roomConfig.HashChecker.SubscribedLists == nil {
		log.Printf("Room %s is not subscribed to any lists!", roomConfig.RoomID)
		return false // Not subscribed to any lists
	}

	subMap := make(map[string]bool)

	for _, listId := range roomConfig.HashChecker.SubscribedLists {
		subMap[listId.Hex()] = true
	}

	found := false
	log.Printf("%v", subMap)

	for _, listId := range hashObj.PartOf {
		_, ok := subMap[listId.Hex()]

		if ok {
			found = true
			break
		}
	}

	if !found {
		log.Printf("Room %s is not subscribed to any lists of hashobj %s!", roomConfig.RoomID, hashObj.ID.Hex())
		return false // Not subscribed
	}

	return true
}

// handleIllegalContent is called when a hash-match is found to take configured actions
func handleIllegalContent(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig, parentCtx context.Context) {
	ctx, span := tracer.Tracer.Start(parentCtx, "handle_hashing")
	defer span.End()

	filesFlagged.Inc()

	switch roomConfig.HashChecker.HashCheckMode {
	case 0:
		postNotice(evt, matrixClient, hashObj, roomConfig, ctx)
		break
	case 1:
		redactMessage(evt, matrixClient, hashObj, ctx)
		if roomConfig.HashChecker.NoticeToChat {
			postNotice(evt, matrixClient, hashObj, roomConfig, ctx)
		}
		break
	case 2:
		muteUser(evt, matrixClient, ctx)
		redactMessage(evt, matrixClient, hashObj, ctx)
		if roomConfig.HashChecker.NoticeToChat {
			postNotice(evt, matrixClient, hashObj, roomConfig, ctx)
		}
		break
	case 3:
		banUser(evt, matrixClient, hashObj, ctx)
		redactMessage(evt, matrixClient, hashObj, ctx)
		if roomConfig.HashChecker.NoticeToChat {
			postNotice(evt, matrixClient, hashObj, roomConfig, ctx)
		}
		break
	}
}

// redactMessage deletes the message sent in the given event
func redactMessage(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, ctx context.Context) {
	ctx, span := tracer.Tracer.Start(ctx, "redact_message")
	defer span.End()

	opts := mautrix.ReqRedact{Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex())}
	_, err := matrixClient.RedactEvent(evt.RoomID, evt.ID, opts)
	if err != nil {
		span.RecordError(err)
		log.Printf("ERROR: Could not redact event - %v", err)
	}
}

// muteUser sets a users power-level to -1 to prevent them from sending messages
func muteUser(evt *event.Event, matrixClient *mautrix.Client, ctx context.Context) {
	ctx, span := tracer.Tracer.Start(ctx, "mute_user")
	defer span.End()

	plEventContent, err := GetRoomPowerLevelState(matrixClient, evt.RoomID)
	if err != nil {
		span.RecordError(err)
		log.Printf("ERROR: Could mute user - %v", err)
		return
	}

	plEventContent.Users[evt.Sender.String()] = -1

	_, err = matrixClient.SendStateEvent(evt.RoomID, event.StatePowerLevels, "", plEventContent)
	if err != nil {
		span.RecordError(err)
		log.Printf("ERROR: Could mute user - %v", err)
		return
	}
}

// banUser bans the sender of an event from the room
func banUser(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, ctx context.Context) {
	ctx, span := tracer.Tracer.Start(ctx, "post_notice")
	defer span.End()

	req := mautrix.ReqBanUser{
		Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex()),
		UserID: evt.Sender,
	}

	_, err := matrixClient.BanUser(evt.RoomID, &req)
	if err != nil {
		span.RecordError(err)
	}
}

// postNotice posts a notice about the given event into its room
func postNotice(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig, ctx context.Context) {
	ctx, span := tracer.Tracer.Start(ctx, "post_notice")
	defer span.End()

	local, server, err := evt.Sender.Parse()
	if err != nil {
		span.RecordError(err)
		return
	}
	_, err = matrixClient.SendNotice(evt.RoomID, fmt.Sprintf(
		`Veles Triggered: The message by %s (on %s) was flagged for containing material used by spammers or trolls!

If you believe this action was an accident, please contact an room administrator or moderator. (Reference: %s)`, local, server, hashObj.ID.Hex()))
	if err != nil {
		span.RecordError(err)
		return
	}

}

/*func msgMods(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig) {
	local, server, err := evt.Sender.Parse()
	if err != nil {
		return
	}
	SendAlert(matrixClient, evt.RoomID.String(), fmt.Sprintf(
		`Veles Triggered: The message by %s (on %s) was flagged for containing material used by spammers or trolls! (Reference: %s)`, local, server, hashObj.ID.Hex()))
}*/
