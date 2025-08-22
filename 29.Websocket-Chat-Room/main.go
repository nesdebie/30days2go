package main

import (
    "fmt"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var mutex = &sync.Mutex{}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer ws.Close()

    clientID := fmt.Sprintf("%p", ws)

    mutex.Lock()
    clients[ws] = true
    mutex.Unlock()

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            mutex.Lock()
            delete(clients, ws)
            mutex.Unlock()
            break
        }

        prefixedMsg := []byte(clientID + ": " + string(msg))

        mutex.Lock()
        for client := range clients {
            if client != ws {
                if err := client.WriteMessage(websocket.TextMessage, prefixedMsg); err != nil {
                    client.Close()
                    delete(clients, client)
                }
            }
        }
        mutex.Unlock()
    }
}


func main() {
    http.HandleFunc("/ws", handleConnections)
    fmt.Println("Serveur WebSocket launched.")
	fmt.Println("Use `websocat ws://localhost:8080/ws` to join.")
    http.ListenAndServe(":8080", nil)
}
