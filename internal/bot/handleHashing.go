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

	defer func(reader io.ReadCloser) {
		_ = reader.Close()
	}(reader)

	hashWriter := sha512.New()
	if _, err = io.Copy(hashWriter, reader); err != nil {
		log.Printf("Error: Could not hash file from Content-URL: \"%s\" - %v", content.URL, err)
		return
	}

	sum := hex.EncodeToString(hashWriter.Sum(nil))

	roomConfig := config.GetRoomConfig(evt.RoomID.String())

	hashObj, err := db.GetEntryByHash(sum)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			/*entry := model.DBEntry{
				ID:        primitive.NewObjectID(),
				HashValue: sum,
				FileURL:   "placeholder",
				Timestamp: time.Now(),
				AddedBy:   nil,
				Comments:  nil,
			}

			db.SaveEntry(&entry)*/

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

		var buf bytes.Buffer

		json.Indent(&buf, jsonVal, "", "  ")

		matrixClient.SendNotice(evt.RoomID, fmt.Sprintf("DEBUG:\n%s", buf.String()))
	}

	if roomConfig.HashChecker.SubscribedLists == nil {
		log.Printf("Room %s is not subscribed to any lists!", roomConfig.RoomID)
		return // Not subscribed to any lists
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
		return // Not subscribed
	}

	log.Printf("Illegal content detected in room %s!", roomConfig.RoomID)

	handleIllegalContent(evt, matrixClient, hashObj, roomConfig)
}

func redactMessage(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry) {
	opts := mautrix.ReqRedact{Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex())}
	_, err := matrixClient.RedactEvent(evt.RoomID, evt.ID, opts)
	if err != nil {
		log.Printf("ERROR: Could not redact event - %v", err)
	}
}

func muteUser(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry) {
	plEventContent, err := GetRoomState(matrixClient, evt.RoomID)
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

func banUser(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry) {
	req := mautrix.ReqBanUser{
		Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex()),
		UserID: evt.Sender,
	}

	matrixClient.BanUser(evt.RoomID, &req)
}

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
		muteUser(evt, matrixClient, hashObj)
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
