package main

import (
	"log"
	"net/http"
	"workout-api/internal/database"
	"workout-api/internal/handlers"
	"workout-api/internal/repository"
	"workout-api/internal/routes"
	"workout-api/internal/services"
)

func main() {
	// Initialize database connection
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repository, service, and handler
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup router with handlers
	r := routes.SetupRouter(userHandler)

	log.Println("Starting server on 8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
