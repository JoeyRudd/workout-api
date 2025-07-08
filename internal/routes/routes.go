package routes

import (
	"github.com/gin-gonic/gin"
	"workout-api/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Router check
	r.GET("/ping", handlers.Ping)

	// User routes
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)

	return r
}
