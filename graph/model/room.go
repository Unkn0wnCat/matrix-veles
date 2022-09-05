package model

import "github.com/Unkn0wnCat/matrix-veles/internal/config"

type Room struct {
	ID                string             `json:"id"`
	Active            bool               `json:"active"`
	Deactivated       bool               `json:"deactivated"`
	Name              string             `json:"name"`
	RoomID            string             `json:"roomId"`
	Debug             bool               `json:"debug"`
	AdminPowerLevel   int                `json:"adminPowerLevel"`
	HashCheckerConfig *HashCheckerConfig `json:"hashCheckerConfig"`
}

type HashCheckerConfig struct {
	ChatNotice      bool            `json:"chatNotice"`
	HashCheckMode   HashCheckerMode `json:"hashCheckMode"`
	SubscribedLists []string        `json:"subscribedLists"`
}

func MakeRoom(room *config.RoomConfig) *Room {
	var subscribed []string

	for _, subId := range room.HashChecker.SubscribedLists {
		subscribed = append(subscribed, subId.Hex())
	}

	return &Room{
		ID:              room.ID.Hex(),
		Name:            room.Name,
		Active:          room.Active && !room.Deactivate,
		Deactivated:     room.Deactivate,
		RoomID:          room.RoomID,
		Debug:           room.Debug,
		AdminPowerLevel: room.AdminPowerLevel,
		HashCheckerConfig: &HashCheckerConfig{
			ChatNotice:      room.HashChecker.NoticeToChat,
			HashCheckMode:   AllHashCheckerMode[room.HashChecker.HashCheckMode],
			SubscribedLists: subscribed,
		},
	}
}
