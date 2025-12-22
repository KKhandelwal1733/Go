package main

import (
	"fmt"
	"net/http"
	"time"

	"example.com/myapp/models"
	"github.com/gin-gonic/gin"
)



func main() {
	server:=gin.Default()
	server.GET("/events",getEvents)
	server.POST("/events",createEvent)


	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events:=models.GetAllEvents()
	c.JSON(http.StatusOK,events)}


func createEvent(c *gin.Context){
   var event models.Event
   err:= c.ShouldBindJSON(&event)
   fmt.Println(event)
   fmt.Println(err)
   if err!=nil{
    c.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse the data you sent!!!"})
    return }
   event.ID=1;
   event.UserID=1;
   event.Date=time.Now()
   event.Save()
   c.JSON(http.StatusCreated,gin.H{"messge":"Event created!","event":event})

   
}



