package routes

import (
	"net/http"
	"strconv"

	"example.com/myapp/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data you sent!!!"})
		return
	}
	userId := c.GetInt64("userId")
	
	event.UserID = userId
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event_id": event.ID})

}

func getEventByID(c *gin.Context) {
	id,err:=strconv.ParseInt(c.Param("id"),10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid event ID"})
		return
	}
	event,err:=models.GetEventByID(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Could not retrieve event"})
		return
	}
	c.JSON(http.StatusOK,*event)

}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event"})
		return
	}
	if event.UserID != c.GetInt64("userId") {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to update this event"})
		return
	}
	err=c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data you sent"})
		return
	}
	event.ID = id

	err=event.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event_id": event.ID})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event"})
		return
	}
	if event.UserID != c.GetInt64("userId") {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not authorized to delete this event"})
		return
	}

	err = models.DeleteEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
