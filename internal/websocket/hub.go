package websocket

import (
    "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func HandleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteMessage(websocket.TextMessage, []byte(msg))
            if err != nil {
                client.Close()
                delete(clients, client)
            }
        }
    }
}
