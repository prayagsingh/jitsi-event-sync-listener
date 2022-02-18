package models

import (
	"github.com/google/uuid"
)

// UserDetails for the handling All Occupants in JSON sent by the muc-room-destroyed event
type UserDetails struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Id          uuid.UUID `json:"id" gorm:"unique; type:uuid; column:id; default:uuid_generate_v4(); not_null"`
	OccupantJid string    `json:"occupant_jid"`
	JoinedAt    int64     `json:"joined_at"`
	LeftAt      int64     `json:"left_at"`
}

// RequestData for the handling the JSON request sent by events
type RoomDestroyed struct {
	EventName    string        `json:"event_name"`
	RoomName     string        `json:"room_name"`
	RoomJID      string        `json:"room_jid" gorm:"primaryKey"`
	StartedAt    int64         `json:"created_at"`
	DestroyedAt  int64         `json:"destroyed_at"`
	AllOccupants []UserDetails `json:"all_occupants"`
}

type RoomEvents struct {
	EventName   string `json:"event_name"`
	RoomName    string `json:"room_name"`
	RoomJID     string `json:"room_jid" gorm:"primaryKey"`
	StartedAt   int64  `json:"created_at"`
	DestroyedAt int64  `json:"destroyed_at"`
	Occupant    User   `json:"occupant"`
}

// User for the handling the JSON request sent by muc-occupant-joined events
type User struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Id          uuid.UUID `json:"id" gorm:"unique; type:uuid; column:id; default:uuid_generate_v4(); not_null"`
	OccupantJid string    `json:"occupant_jid"`
	JoinedAt    int64     `json:"joined_at"`
	LeftAt      int64     `json:"left_at"`
}

// SaveRoomDestroyed for sending data to DB
type SaveRoomDestroyed struct {
	EventName    string
	RoomName     string
	RoomJID      string
	StartedAt    int64
	DestroyedAt  int64
	AllOccupants []UserDetails
}

// SaveRoomEvents for sending data to DB
type SaveRoomEvents struct {
	EventName   string
	RoomName    string
	RoomJID     string
	StartedAt   int64
	DestroyedAt int64
	Name        string
	Email       string
	Id          uuid.UUID
	OccupantJid string
	JoinedAt    int64
	LeftAt      int64
}
