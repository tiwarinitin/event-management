package routes

import (
	"net/http"
	"strconv"

	"example.com/event-mgmt/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.Cancel(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel event!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Cancelled!"})
}
