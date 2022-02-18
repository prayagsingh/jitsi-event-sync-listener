package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayagsingh/jitsi-event-sync-listner/connection"
	"github.com/prayagsingh/jitsi-event-sync-listner/models"
)

// RoomCreated handles the room_created event
func RoomCreated(c *gin.Context) {

	log.Println("room created event received")

	db := connection.GetPostgresDB()

	log.Println("value of db is: ", db)
	// initializing the request data struct
	req_data := models.RoomEvents{}

	// fetching the data from the request
	if err := c.ShouldBindJSON(&req_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveData := models.SaveRoomEvents{
		EventName:   req_data.EventName,
		RoomName:    req_data.RoomName,
		RoomJID:     req_data.RoomJID,
		StartedAt:   req_data.StartedAt,
		DestroyedAt: req_data.DestroyedAt,
		Name:        req_data.Occupant.Name,
		Email:       req_data.Occupant.Email,
		Id:          req_data.Occupant.Id,
		OccupantJid: req_data.Occupant.OccupantJid,
		JoinedAt:    req_data.Occupant.JoinedAt,
		LeftAt:      req_data.Occupant.LeftAt,
	}

	db.Create(&saveData)

	c.JSON(http.StatusOK, gin.H{"room-created": saveData})
}

func OccupantJoined(c *gin.Context) {

	log.Println("occupant joined event received")

	db := connection.GetPostgresDB()

	// initializing the request data struct
	req_data := models.RoomEvents{}

	// fetching the data from the request
	if err := c.ShouldBindJSON(&req_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveData := models.SaveRoomEvents{
		EventName:   req_data.EventName,
		RoomName:    req_data.RoomName,
		RoomJID:     req_data.RoomJID,
		StartedAt:   req_data.StartedAt,
		DestroyedAt: req_data.DestroyedAt,
		Name:        req_data.Occupant.Name,
		Email:       req_data.Occupant.Email,
		Id:          req_data.Occupant.Id,
		OccupantJid: req_data.Occupant.OccupantJid,
		JoinedAt:    req_data.Occupant.JoinedAt,
		LeftAt:      req_data.Occupant.LeftAt,
	}

	db.Create(&saveData)

	c.JSON(http.StatusOK, gin.H{"occupant-joined": saveData})
}

func OccupantLeft(c *gin.Context) {

	log.Println("occupant left event received")

	db := connection.GetPostgresDB()

	// initializing the request data struct
	req_data := models.RoomEvents{}

	// fetching the data from the request
	if err := c.ShouldBindJSON(&req_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveData := models.SaveRoomEvents{
		EventName:   req_data.EventName,
		RoomName:    req_data.RoomName,
		RoomJID:     req_data.RoomJID,
		StartedAt:   req_data.StartedAt,
		DestroyedAt: req_data.DestroyedAt,
		Name:        req_data.Occupant.Name,
		Email:       req_data.Occupant.Email,
		Id:          req_data.Occupant.Id,
		OccupantJid: req_data.Occupant.OccupantJid,
		JoinedAt:    req_data.Occupant.JoinedAt,
		LeftAt:      req_data.Occupant.LeftAt,
	}

	db.Create(&saveData)

	c.JSON(http.StatusOK, gin.H{"occupant-left": saveData})
}

func RoomDestroyed(c *gin.Context) {
	log.Println("room destroyed event received")

	db := connection.GetPostgresDB()
	// initializing the request data struct
	req_data := models.RoomDestroyed{}

	// fetching the data from the request
	if err := c.ShouldBindJSON(&req_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveData := models.SaveRoomDestroyed{
		EventName:   req_data.EventName,
		RoomName:    req_data.RoomName,
		RoomJID:     req_data.RoomJID,
		StartedAt:   req_data.StartedAt,
		DestroyedAt: req_data.DestroyedAt,
		AllOccupants: req_data.AllOccupants,
	}

	db.Create(&saveData)

	c.JSON(http.StatusOK, gin.H{"room-destroyed": saveData})
}
