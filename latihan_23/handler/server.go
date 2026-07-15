package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ChatServer struct {
	clients clientList
	sync.RWMutex
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(clientList),
	}
}

func (cs *ChatServer) ServeConnection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !websocket.IsWebSocketUpgrade(r) {
		http.Error(w, "websocket upgrade required", http.StatusBadRequest)
		return
	}

	log.Println("New connection")
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade failed:", err)
		return
	}

	client := NewClient(conn, cs)
	cs.addClient(client)

	go client.readMessages()
	go client.writeMessages()

}

func (cs *ChatServer) addClient(client *Client) {

	cs.Lock()
	defer cs.Unlock()

	if _, ok := cs.clients[client]; !ok {

		cs.clients[client] = true
	}

}

func (cs *ChatServer) removeClient(client *Client) {

	cs.Lock()
	defer cs.Unlock()

	if _, ok := cs.clients[client]; ok {

		client.connection.Close()
		delete(cs.clients, client)
	}

}
