package bot

import (
	"encoding/json"
	"fmt"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type StateEventPowerLevel struct {
	Type           string                      `json:"type"`
	Sender         string                      `json:"sender"`
	RoomID         string                      `json:"room_id"`
	EventID        string                      `json:"event_id"`
	OriginServerTS int64                       `json:"origin_server_ts"`
	Content        StateEventPowerLevelContent `json:"content"`
	Unsigned       struct {
		Age int `json:"age"`
	} `json:"unsigned"`
}

type StateEventPowerLevelContent struct {
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

// GetRoomPowerLevelState returns the rooms current power levels from the state
func GetRoomPowerLevelState(matrixClient *mautrix.Client, roomId id.RoomID) (*StateEventPowerLevelContent, error) {
	// https://matrix.example.com/_matrix/client/r0/rooms/<roomId.String()>/state
	url := matrixClient.BuildURL("rooms", roomId.String(), "state")

	res, err := matrixClient.MakeRequest("GET", url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Could request room state - %v", err)
	}

	// res contains an array of state events
	var stateEvents []StateEventPowerLevel

	err = json.Unmarshal(res, &stateEvents)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Could parse room state - %v", err)
	}

	// plEventContent will hold the final event
	var plEventContent StateEventPowerLevelContent

	found := false

	for _, e2 := range stateEvents {
		if e2.Type != event.StatePowerLevels.Type {
			continue // If the current event is not of the power level, skip.
		}

		// This is what we're looking for!
		found = true
		plEventContent = e2.Content
		break
	}

	if !found {
		return nil, fmt.Errorf("ERROR: Could find room power level - %v", err)
	}

	// The following handle cases in which empty lists may have been parsed as nil
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
