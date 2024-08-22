package main

import (
	"example.com/event-mgmt/db"
	"example.com/event-mgmt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
