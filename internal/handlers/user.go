package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workout-api/internal/models"
)

// Temporary in memory storage
var users = make(map[string]models.User)

func CreateUser(c *gin.Context) {
	var user models.User

	// Bind and validate JSON request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check id the user already exits
	if _, exists := users[user.ID]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	// Store the user
	users[user.ID] = user

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func GetUser(c *gin.Context) {
	// get user
	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
