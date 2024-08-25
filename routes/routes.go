package routes

import (
	"example.com/event-mgmt/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events", deleteAllEvents)
	authenticated.DELETE("/events/:id", deleteEventById)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
