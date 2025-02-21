package routes

import (
    "github.com/gorilla/mux"
    "github.com/Dharsansivaloganathan/task-manager/internal/handlers"
    "github.com/Dharsansivaloganathan/task-manager/internal/websocket"

)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    

    // Authentication routes
    router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
    router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

    // Task routes
    router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")

    // WebSocket route
    router.HandleFunc("/ws", websocket.HandleConnections)

    return router
}
