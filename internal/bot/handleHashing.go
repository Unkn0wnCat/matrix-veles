package bot

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Unkn0wnCat/matrix-veles/internal/config"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
)

// handleHashing hashes and checks a message, taking configured actions on match
func handleHashing(content *event.MessageEventContent, evt *event.Event, matrixClient *mautrix.Client) {
	url, err := content.URL.Parse()
	if err != nil {
		log.Printf("Error: Could not parse Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	reader, err := matrixClient.Download(url)
	if err != nil {
		log.Printf("Error: Could not read file from Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	defer func(reader io.ReadCloser) { _ = reader.Close() }(reader)

	hashWriter := sha512.New()
	if _, err = io.Copy(hashWriter, reader); err != nil {
		log.Printf("Error: Could not hash file from Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	sum := hex.EncodeToString(hashWriter.Sum(nil))

	// Fetch room configuration for adjusting behaviour
	roomConfig := config.GetRoomConfig(evt.RoomID.String())

	hashObj, err := db.GetEntryByHash(sum)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			if roomConfig.Debug {
				matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG - This file is not on the hashlist: %s", sum))
			}
			return
		}
		if roomConfig.Debug {
			matrixClient.SendNotice(evt.RoomID, "DEBUG - Failed to check file. See log.")
		}
		fmt.Printf("Error trying to check database: %v", err)
		return
	}

	if roomConfig.Debug {
		matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG !!! This file is on the hashlist: %s", sum))
		jsonVal, _ := json.Marshal(hashObj)
		matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG:\n%s", makeFancyJSON(jsonVal)))
	}

	if !checkSubscription(&roomConfig, hashObj) {
		return
	}

	log.Printf("Illegal content detected in room %s!", roomConfig.RoomID)

	handleIllegalContent(evt, matrixClient, hashObj, roomConfig)
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
func handleIllegalContent(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig) {
	switch roomConfig.HashChecker.HashCheckMode {
	case 0:
		postNotice(evt, matrixClient, hashObj, roomConfig)
		break
	case 1:
		redactMessage(evt, matrixClient, hashObj)
		if roomConfig.HashChecker.NoticeToChat {
			postNotice(evt, matrixClient, hashObj, roomConfig)
		}
		break
	case 2:
		muteUser(evt, matrixClient)
		redactMessage(evt, matrixClient, hashObj)
		if roomConfig.HashChecker.NoticeToChat {
			postNotice(evt, matrixClient, hashObj, roomConfig)
		}
		break
	case 3:
		banUser(evt, matrixClient, hashObj)
		redactMessage(evt, matrixClient, hashObj)
		if roomConfig.HashChecker.NoticeToChat {
			postNotice(evt, matrixClient, hashObj, roomConfig)
		}
		break
	}
}

// redactMessage deletes the message sent in the given event
func redactMessage(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry) {
	opts := mautrix.ReqRedact{Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex())}
	_, err := matrixClient.RedactEvent(evt.RoomID, evt.ID, opts)
	if err != nil {
		log.Printf("ERROR: Could not redact event - %v", err)
	}
}

// muteUser sets a users power-level to -1 to prevent them from sending messages
func muteUser(evt *event.Event, matrixClient *mautrix.Client) {
	plEventContent, err := GetRoomPowerLevelState(matrixClient, evt.RoomID)
	if err != nil {
		log.Printf("ERROR: Could mute user - %v", err)
		return
	}

	plEventContent.Users[evt.Sender.String()] = -1

	_, err = matrixClient.SendStateEvent(evt.RoomID, event.StatePowerLevels, "", plEventContent)
	if err != nil {
		log.Printf("ERROR: Could mute user - %v", err)
		return
	}
}

// banUser bans the sender of an event from the room
func banUser(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry) {
	req := mautrix.ReqBanUser{
		Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex()),
		UserID: evt.Sender,
	}

	matrixClient.BanUser(evt.RoomID, &req)
}

// postNotice posts a notice about the given event into its room
func postNotice(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig) {
	local, server, err := evt.Sender.Parse()
	if err != nil {
		return
	}
	matrixClient.SendNotice(evt.RoomID, fmt.Sprintf(
		`Veles Triggered: The message by %s (on %s) was flagged for containing material used by spammers or trolls!

If you believe this action was an accident, please contact an room administrator or moderator. (Reference: %s)`, local, server, hashObj.ID.Hex()))

}

/*func msgMods(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig) {
	local, server, err := evt.Sender.Parse()
	if err != nil {
		return
	}
	SendAlert(matrixClient, evt.RoomID.String(), fmt.Sprintf(
		`Veles Triggered: The message by %s (on %s) was flagged for containing material used by spammers or trolls! (Reference: %s)`, local, server, hashObj.ID.Hex()))
}*/
