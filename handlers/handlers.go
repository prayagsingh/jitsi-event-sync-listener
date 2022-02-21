package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayagsingh/jitsi-event-sync-listener/connection"
	"github.com/prayagsingh/jitsi-event-sync-listener/models"
)

// RoomCreated handles the room_created event
func RoomCreated(c *gin.Context) {

	log.Println("room created event received")

	db := connection.GetPostgresDB()

	//log.Println("value of db is: ", db)
	// initializing the request data struct
	req_data := models.RoomEvents{}

	// fetching the data from the request
	if err := c.ShouldBindJSON(&req_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&req_data)

	c.JSON(http.StatusOK, gin.H{"room-created": req_data})
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

	//fmt.Println("value of req_data is: ", req_data)
	db.Create(&req_data)

	c.JSON(http.StatusOK, gin.H{"occupant-joined": req_data})
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

	db.Create(&req_data)

	c.JSON(http.StatusOK, gin.H{"occupant-left": req_data})
}

func RoomDestroyed(c *gin.Context) {
	log.Println("room destroyed event received")

	db := connection.GetPostgresDB()
	// initializing the request data struct
	req_data := models.RoomDestroyedEvents{}

	// fetching the data from the request
	if err := c.ShouldBindJSON(&req_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&req_data)

	c.JSON(http.StatusOK, gin.H{"room-destroyed": req_data})
}
