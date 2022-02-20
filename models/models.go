package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoomDestroyedEvent is the struct for the room_destroyed event
type RoomDestroyedEvents struct {
	gorm.Model
	EventName    string                     `json:"event_name"`
	RoomName     string                     `json:"room_name"`
	RoomJID      string                     `json:"room_jid"`
	StartedAt    int64                      `json:"created_at"`
	DestroyedAt  int64                      `json:"destroyed_at"`
	AllOccupants []RoomDestroyedUserDetails `json:"all_occupants"`
}

// RoomDestroyedUserDetails for the handling the JSON request sent by muc-occupant-joined events
type RoomDestroyedUserDetails struct {
	gorm.Model                   // need to increment the ID to avoid the conflict
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	UserId                string `json:"id"`
	OccupantJid           string `json:"occupant_jid"`
	JoinedAt              int64  `json:"joined_at"`
	LeftAt                int64  `json:"left_at"`
	RoomDestroyedEventsID uint   // mandatory for creating relationship between the two tables
}

// RoomEvents is the struct for the room_created event
type RoomEvents struct {
	gorm.Model
	EventName   string      `json:"event_name"`
	RoomName    string      `json:"room_name"`
	RoomJID     string      `json:"room_jid"`
	StartedAt   int64       `json:"created_at"`
	DestroyedAt int64       `json:"destroyed_at"`
	Occupant    UserDetails `gorm:"embedded" json:"occupant"` // gorm embedded the UserDetails struct into RoomEvents
}

// UserDetails for the handling the JSON request sent by muc-occupant-joined events
type UserDetails struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	UserId      string `json:"id"`
	OccupantJid string `json:"occupant_jid"`
	JoinedAt    int64  `json:"joined_at"`
	LeftAt      int64  `json:"left_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *UserDetails) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.UserId = uuid.NewString()
	return
}

func (user *RoomDestroyedUserDetails) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.UserId = uuid.NewString()
	return
}
