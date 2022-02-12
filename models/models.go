package models

type userDetails struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Id          string `json:"id"`
	OccupantJid string `json:"occupant_jid"`
	JoinedAt    int64  `json:"joined_at"`
	LeftAt      int64  `json:"left_at"`
}

type RequestData struct {
	EventName    string        `json:"event_name"`
	RoomName     string        `json:"room_name"`
	RoomJID      string        `json:"room_jid"`
	CreatedAt    int64         `json:"created_at"`
	DestroyedAt  int64         `json:"destroyed_at"`
	AllOccupants []userDetails `json:"all_occupants"`
}
