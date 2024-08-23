package routes

import (
	"net/http"

	"example.com/event-mgmt/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.Users

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
