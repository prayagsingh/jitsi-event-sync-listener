package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayagsingh/jitsi-event-sync-listener/connection"
	"github.com/prayagsingh/jitsi-event-sync-listener/routes"
	"github.com/spf13/viper"
)

var PORT = ":7002"

// creating a golang http server
func main() {

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	db_user := viper.Get("POSTGRES_USER")
	db_password := viper.Get("POSTGRES_PASSWORD")
	db_name := viper.Get("POSTGRES_DATABASE")
	db_host := viper.Get("POSTGRES_HOST")
	db_port := viper.Get("POSTGRES_PORT")

	log.Printf("Connecting to DB: %v:%v@%v:%v/%v", db_user, db_password, db_host, db_port, db_name)
	// Setup the DB
	connection.SetupDB(db_host.(string), db_user.(string), db_name.(string), db_password.(string), db_port.(string))

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	routes.Init(router)

	srv := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Error starting or closing listener
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal from the OS.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
