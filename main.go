package main

import (
	"net/http"
	"strconv"

	"example.com/myapp/db"
	"example.com/myapp/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/event/:id",getEventByID)

	server.Run(":8080")
}

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

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully", "event_id": event.ID})

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
