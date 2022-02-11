package models

import "time"

type userDetails struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Id          string `json:"id"`
	OccupantJid string `json:"occupant_jid"`
	JoinedAt    int    `json:"joined_at"`
	LeftAt      int    `json:"left_at"`
}

type RequestData struct {
	EventName    string        `json:"event_name"`
	RoomName     string        `json:"room_name"`
	RoomJID      string        `json:"room_jid"`
	CreatedAt    time.Duration `json:"created_at"`
	DestroyedAt  time.Duration `json:"destroyed_at"`
	AllOccupants []userDetails `json:"all_occupants"`
}
