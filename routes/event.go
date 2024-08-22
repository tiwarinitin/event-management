package routes

import (
	"net/http"
	"strconv"

	"example.com/event-mgmt/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event. Try again"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data!"})
		return
	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event. Try again"})
		return
	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not bind data!"})
		return
	}

	updateEvent.ID = eventID
	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated!"})
}

func deleteEventById(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event. Try again"})
		return
	}

	var deleteEvent models.Event
	deleteEvent.ID = eventID
	err = deleteEvent.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted!"})
}

func deleteAllEvents(context *gin.Context) {
	var deleteEvent models.Event
	err := deleteEvent.DeleteAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "All Event Deleted!"})
}
