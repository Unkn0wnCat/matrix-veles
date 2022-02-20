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
	"maunium.net/go/mautrix/id"
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

	handleIllegalContent(evt, matrixClient, hashObj, roomConfig)
}

func redactMessage(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry) {
	opts := mautrix.ReqRedact{Reason: fmt.Sprintf("Veles has detected an hash-map-match! Tags: %s, ID: %s", hashObj.Tags, hashObj.ID.Hex())}
	_, err := matrixClient.RedactEvent(evt.RoomID, evt.ID, opts)
	if err != nil {
		log.Printf("ERROR: Could not redact event - %v", err)
	}
}

type StateEventPL struct {
	Type           string              `json:"type"`
	Sender         string              `json:"sender"`
	RoomID         string              `json:"room_id"`
	EventID        string              `json:"event_id"`
	OriginServerTS int64               `json:"origin_server_ts"`
	Content        StateEventPLContent `json:"content"`
	Unsigned       struct {
		Age int `json:"age"`
	} `json:"unsigned"`
}

type StateEventPLContent struct {
	Ban           int            `json:"ban"`
	Events        map[string]int `json:"events"`
	EventsDefault int            `json:"events_default"`
	Invite        int            `json:"invite"`
	Kick          int            `json:"kick"`
	Notifications map[string]int `json:"notifications"`
	Redact        int            `json:"redact"`
	StateDefault  int            `json:"state_default"`
	Users         map[string]int `json:"users"`
	UsersDefault  int            `json:"users_default"`
}

func GetRoomState(matrixClient *mautrix.Client, roomId id.RoomID) (*StateEventPLContent, error) {
	url := matrixClient.BuildURL("rooms", roomId.String(), "state")
	log.Println(url)
	res, err := matrixClient.MakeRequest("GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Could request room state - %v", err)
	}

	var stateEvents []StateEventPL

	err = json.Unmarshal(res, &stateEvents)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Could parse room state - %v", err)
	}

	var plEventContent StateEventPLContent

	found := false

	for _, e2 := range stateEvents {
		if e2.Type != event.StatePowerLevels.Type {
			continue
		}

		found = true
		plEventContent = e2.Content
	}

	if !found {
		return nil, fmt.Errorf("ERROR: Could find room power level - %v", err)
	}

	if plEventContent.Events == nil {
		plEventContent.Events = make(map[string]int)
	}

	if plEventContent.Notifications == nil {
		plEventContent.Notifications = make(map[string]int)
	}

	if plEventContent.Users == nil {
		plEventContent.Users = make(map[string]int)
	}

	return &plEventContent, nil
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

func postNotice(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig) {
	local, server, err := evt.Sender.Parse()
	if err != nil {
		return
	}
	matrixClient.SendNotice(evt.RoomID, fmt.Sprintf(
		`Veles Triggered: The message by %s (on %s) was flagged for containing material used by spammers or trolls!

If you believe this action was an accident, please contact an room administrator or moderator. (Reference: %s)`, local, server, hashObj.ID.Hex()))

}

func handleIllegalContent(evt *event.Event, matrixClient *mautrix.Client, hashObj *model.DBEntry, roomConfig config.RoomConfig) {
	switch roomConfig.HashCheckMode {
	case 0:
		break
	case 1:
		break
	case 2:
		muteUser(evt, matrixClient, hashObj)
		redactMessage(evt, matrixClient, hashObj)
		postNotice(evt, matrixClient, hashObj, roomConfig)
		break
	case 3:
		break

	}
}
