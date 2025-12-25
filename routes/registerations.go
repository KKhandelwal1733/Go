package routes

import (
	"net/http"
	"strconv"
	"example.com/myapp/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event"})
		return
	}
	err = models.RegisterUserForEvent(userId, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func unregisterFromEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	err = models.UnregisterUserFromEvent(userId, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event", "error": err.Error()})
		return
	}		
	c.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})
}

func getUserRegistrations(c *gin.Context) {
	userId := c.GetInt64("userId")
	events, err := models.GetEventsForUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve registrations", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}