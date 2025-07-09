package routes

import (
	"github.com/gin-gonic/gin"
	"workout-api/internal/handlers"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	// Router check
	r.GET("/ping", handlers.Ping)

	// User routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.GET("/users", userHandler.GetAllUsers)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	return r
}
