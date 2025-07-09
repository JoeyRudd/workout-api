package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"workout-api/internal/models"
	"workout-api/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	// Bind and validate JSON request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user using service
	if err := h.userService.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send JSON response status Created
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// get user ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Send JSON response status OK
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	// Send JSON response status OK
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Get user ID from the URL path parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	// Send JSON response status OK
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
