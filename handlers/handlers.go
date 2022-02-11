package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prayagsingh/jitsi-event-sync-listner/models"
)

// RoomCreated handles the room_created event
func RoomCreated(w http.ResponseWriter, r *http.Request) {

	log.Println("room created event received")
	// initializing the request data struct
	req_data := models.RequestData{}

	// fetching the data from the request
	err := json.NewDecoder(r.Body).Decode(&req_data)
	if err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("room-created: ", req_data)
}

func OccupantJoined(w http.ResponseWriter, r *http.Request) {

	log.Println("occupant joined event received")
	// initializing the request data struct
	req_data := models.RequestData{}

	// fetching the data from the request
	err := json.NewDecoder(r.Body).Decode(&req_data)
	if err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("occupant joined: ", req_data)
}

func OccupantLeft(w http.ResponseWriter, r *http.Request) {

	log.Println("occupant left event received")
	// initializing the request data struct
	req_data := models.RequestData{}

	// fetching the data from the request
	err := json.NewDecoder(r.Body).Decode(&req_data)
	if err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("occupant left: ", req_data)
}

func RoomDestroyed(w http.ResponseWriter, r *http.Request) {

	log.Println("room destroyed event received")
	// initializing the request data struct
	req_data := models.RequestData{}

	// fetching the data from the request
	err := json.NewDecoder(r.Body).Decode(&req_data)
	if err != nil {
		log.Printf("json decoding error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("room destroyed: ", req_data)
}
