package handler

import (
	"chitchater/auth"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin:     checkOrigin,
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

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	allowedOrigins := map[string]bool{
		"http://localhost:8080":  true,
		"http://127.0.0.1:8080":  true,
		"https://localhost:8080": true,
	}

	return allowedOrigins[origin]
}

type ChatServer struct {
	clients  clientList
	handlers map[string]EventHandler
	otps     auth.RetentionMap
	sync.RWMutex
}

func NewChatServer(ctx context.Context) *ChatServer {
	serve := &ChatServer{
		clients:  make(clientList),
		handlers: make(map[string]EventHandler),
		otps:     auth.NewRetentionMap(ctx, 5*time.Second),
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

func (cs *ChatServer) checkOTP(w http.ResponseWriter, r *http.Request) bool {
	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	if !cs.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
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

	if !cs.checkOTP(w, r) {
		log.Println("WebSocket connection rejected: invalid OTP")
		return
	}

	log.Println("New connection")
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Connection closed. ", err)
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

func (cs *ChatServer) LoginHandler(w http.ResponseWriter, r *http.Request) {

	reqAccount := new(auth.Account)

	if err := json.NewDecoder(r.Body).Decode(reqAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc, err := reqAccount.FindAccount(reqAccount.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), reqAccount, reqAccount.Password)
	err = acc.ValidatePassword(ctx)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	type Response struct {
		OTP string `json:"otp"`
	}

	otp := cs.otps.NewOTP()
	var response = Response{
		OTP: otp.Key,
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (cs *ChatServer) SignUpHandler(w http.ResponseWriter, r *http.Request) {

	type RequestAccount struct {
		username string
		password string
	}
	reqAccount := &RequestAccount{}

	if err := json.NewDecoder(r.Body).Decode(reqAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAccount := auth.CreateAccount(reqAccount.username, reqAccount.password)
	if err := newAccount.SaveAccount(); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	w.WriteHeader(http.StatusOK)
}
