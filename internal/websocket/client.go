package websocket

import (
    "fmt"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("WebSocket upgrade failed:", err)
        return
    }
    defer conn.Close()

    clients[conn] = true

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            delete(clients, conn)
            break
        }
        broadcast <- string(msg)
    }
}
