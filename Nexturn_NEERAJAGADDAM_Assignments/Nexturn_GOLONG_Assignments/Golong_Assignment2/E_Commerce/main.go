package main

import (
	"ecommerce-inventory-service/config"
	"ecommerce-inventory-service/handlers"
	"ecommerce-inventory-service/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Create a new router
	router := mux.NewRouter()

	// Register middleware
	router.Use(handlers.LoggingMiddleware)
	router.Use(handlers.AuthMiddleware)

	// Register routes
	routes.RegisterProductRoutes(router)

	// Start the server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
