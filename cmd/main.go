package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dharsansivaloganathan/task-manager/internal/config"
	"github.com/Dharsansivaloganathan/task-manager/internal/database"
	"github.com/Dharsansivaloganathan/task-manager/internal/routes"
	"github.com/Dharsansivaloganathan/task-manager/internal/websocket"
	"github.com/rs/cors"
)

func main() {
        // Load configuration
        config.LoadEnv()

        // Connect to the database
        database.Connect()

        // Start WebSocket message handling
        go websocket.HandleMessages()

        // Define API routes
        router := routes.SetupRoutes()

        // Configure CORS
        c := cors.New(cors.Options{
                AllowedOrigins:   []string{"*"}, // Replace with your frontend origin in production
                AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
                AllowedHeaders:   []string{"Authorization", "Content-Type"},
                AllowCredentials: true,
        })

        // Wrap the router with the CORS middleware
        handler := c.Handler(router)

        // Start the server
        fmt.Println("Server is running on port 8080...")
        log.Fatal(http.ListenAndServe(":8080", handler)) //use the handler
}