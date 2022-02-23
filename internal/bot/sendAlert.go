package bot

/*func SendAlert(matrixClient *mautrix.Client, room string, message string) {
	roomConfig, err := config.GetRoomConfigByRoomID(room)
	if err != nil {
		log.Printf("Failed to get room config - %v", err)
		return
	}

	if roomConfig.AlertChannel == nil {
		roomPLState, err := GetRoomState(matrixClient, id.RoomID(room))
		if err != nil {
			log.Printf("Failed to get room power levels - %v", err)
			return
		}

		var mods []id.UserID

		for member, level := range roomPLState.Users {
			if level >= roomConfig.HashChecker.NotificationPowerLevel {
				mods = append(mods, id.UserID(member))
			}
		}

		req := mautrix.ReqCreateRoom{
			Name:       "Veles Alert Channel",
			Topic:      "Veles Alerts",
			Invite:     mods,
			IsDirect:   true,
			Visibility: "private",
		}

		resp, err := matrixClient.CreateRoom(&req)
		if err != nil {
			log.Printf("Failed to create alert room - %v", err)
			return
		}

		str := resp.RoomID.String()

		roomConfig.AlertChannel = &str

		config.SaveRoomConfig(roomConfig)
	}

	matrixClient.SendNotice(id.RoomID(*roomConfig.AlertChannel), message)
}*/
