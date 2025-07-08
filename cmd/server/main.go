package main

import (
	"log"
	"net/http"
	"workout-api/internal/routes"
)

var db = make(map[string]string)

func main() {
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080

	log.Println("Starting server on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
