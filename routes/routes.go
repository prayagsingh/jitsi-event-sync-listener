package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayagsingh/jitsi-event-sync-listener/handlers"
)

// Init for initializing the routes
func Init(router *gin.Engine) {
	// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix
	v1 := router.Group("/api/v1")
	{
		v1.POST("/events/room/created", handlers.RoomCreated)
		v1.POST("/events/room/destroyed", handlers.RoomDestroyed)
		v1.POST("/events/occupant/joined", handlers.OccupantJoined)
		v1.POST("/events/occupant/left", handlers.OccupantLeft)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run(port)
}
