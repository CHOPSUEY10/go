package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// implementing EventHandler function
func SendMessage(event Event, c *Client) error {

	fmt.Printf("Menerima pesan baru : %s \n", string(event.Payload))
	for client := range c.server.clients {
		client.egress <- event
	}

	return nil

}

type ChatServer struct {
	clients  clientList
	handlers map[string]EventHandler
	sync.RWMutex
}

func NewChatServer() *ChatServer {
	serve := &ChatServer{
		clients:  make(clientList),
		handlers: make(map[string]EventHandler),
	}

	serve.SetupEventHandlers()
	return serve

}

// Initialize eventhandler
func (cs *ChatServer) SetupEventHandlers() {

	cs.handlers[EventSendMessage] = SendMessage

}

func (cs *ChatServer) routeEvent(event Event, c *Client) error {
	// Check if the event type is part of handlers value
	if handler, ok := cs.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
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
