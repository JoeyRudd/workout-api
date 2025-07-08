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

	// Send JSON response status Created
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

	// Send JSON response status OK
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetAllUsers(c *gin.Context) {
	// Create a slice of user structs
	userList := make([]models.User, 0, len(users))
	// Loop over and append each user to userList
	for _, user := range users {
		userList = append(userList, user)
	}

	// Send JSON response status OK
	c.JSON(http.StatusOK, gin.H{"users": userList})
}

func DeleteUser(c *gin.Context) {
	// Get user ID from the URL path parameter
	id := c.Param("id")
	_, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Delte user with specified id from users
	delete(users, id)
	// Send JSON response status OK
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
