package routes

import "github.com/gin-gonic/gin"

func RegisterEventRoutes(router *gin.Engine) {

	router.GET("/events", getEvents)
	router.POST("/events", createEvent)
	router.GET("/event/:id", getEventByID)
	router.PUT("/event/:id", updateEvent)
	router.DELETE("/event/:id", deleteEvent)
}