package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prayagsingh/jitsi-event-sync-listner/handlers"
)

// creating a golang http server
func main() {

	// initialzing logger and printing logs to terminal
	log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

	http.HandleFunc("/api/events/room/created", handlers.RoomCreated)
	http.HandleFunc("/api/events/occupant/joined", handlers.OccupantJoined)
	http.HandleFunc("/api/events/occupant/left", handlers.OccupantLeft)
	http.HandleFunc("/api/events/room/destroyed", handlers.RoomDestroyed)

	// initializng error logger
	// Lshotfile will give the info about the error
	//log.New(os.Stdout, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)

	log.Println("Starting the server on port 7002")
	err := http.ListenAndServe(":7002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
