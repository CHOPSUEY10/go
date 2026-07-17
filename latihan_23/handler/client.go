package handler

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type clientList map[*Client]bool

var typeMsg = map[int]string{
	1:  "TextMessage",
	2:  "BinaryMessage",
	8:  "CloseMessage",
	9:  "PingMessage",
	10: "PongMessage",
}

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

func (c *Client) readMessages() {

	defer func() {
		c.server.removeClient(c)

	}()

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

	}
}

func (c *Client) writeMessages() {

	defer func() {

		c.server.removeClient(c)

	}()

	//writer := bufio.NewWriter(c.connection.NetConn())

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
			}

		}

	}
}
