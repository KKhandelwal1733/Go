package routes

import (
	"example.com/myapp/utils/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(router *gin.Engine) {

	router.GET("/events", getEvents)
	router.GET("/event/:id",getEventByID)	
	router.POST("/signup",signup)
	router.POST("/login",login)
	authenticatedRoutes:=router.Group("/")
	authenticatedRoutes.Use(middleware.AuthMiddleware)
	authenticatedRoutes.POST("/event", createEvent)
	authenticatedRoutes.DELETE("/event/:id", deleteEvent)
	authenticatedRoutes.PUT("/event/:id", updateEvent)
	authenticatedRoutes.POST("/event/:id/register", registerForEvent)
	authenticatedRoutes.GET("/myregistrations", getUserRegistrations)
	authenticatedRoutes.DELETE("/event/:id/unregister", unregisterFromEvent)

}