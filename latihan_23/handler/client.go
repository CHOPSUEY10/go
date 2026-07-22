package handler

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type clientList map[*Client]bool

var (
	typeMsg = map[int]string{
		1:  "TextMessage",
		2:  "BinaryMessage",
		8:  "CloseMessage",
		9:  "PingMessage",
		10: "PongMessage",
	}

	pongWait     = time.Second * 10
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	connection *websocket.Conn
	server     *ChatServer
	egress     chan Event
	//sync.RWMutex
}

func NewClient(conn *websocket.Conn, server *ChatServer) *Client {

	return &Client{

		connection: conn,
		server:     server,
		egress:     make(chan Event),
	}

}

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("Pong...")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

func (c *Client) readMessages() {

	defer func() {
		c.server.removeClient(c)

	}()

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	// Konfigurasi apa yang akan dilakukan client handler setelah mendapatkan pong
	c.connection.SetPongHandler(c.pongHandler)
	// Membatasi ukuran frame yang masuk
	c.connection.SetReadLimit(512)

	for {

		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			// mencetak log error jika error bukan dari client
			if websocket.IsUnexpectedCloseError(err, 1001, 1006, 1003) {
				log.Printf("error reading message : %v", err)
			}

			break
		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error unmashalling event %v", err)
			break
		}

		if err := c.server.routeEvent(request, c); err != nil {
			log.Printf("error handling messsage : %v", err)

		}

		c.connection.SetReadDeadline(time.Now().Add(pongWait))

	}
}

func (c *Client) writeMessages() {

	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.server.removeClient(c)
	}()

	for {
		select {
		case message, filled := <-c.egress:
			if !filled {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("Connection closed : %v", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("Failed to send message %v", err)
				return
			}

		case <-ticker.C:
			log.Println("Ping...")

			if err := c.connection.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Printf("Connection closed")
				return
			}
		}
	}
}
