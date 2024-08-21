package main

import (
	"net/http"

	"example.com/event-mgmt/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	events := models.GetAllEvent()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}
