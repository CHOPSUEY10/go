package handler

import (
	"log"

	"github.com/gorilla/websocket"
)

type clientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	server     *ChatServer
	egress     chan []byte
}

func NewClient(conn *websocket.Conn, server *ChatServer) *Client {

	return &Client{

		connection: conn,
		server:     server,
		egress:     make(chan []byte),
	}

}

var typeMsg = map[int]string{
	1: "TextMessage",

	2: "BinaryMessage",

	8: "CloseMessage",

	9: "PingMessage",

	10: "PongMessage",
}

func (c *Client) readMessages() {

	defer func() {
		c.server.removeClient(c)

	}()

	for {

		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			// mencetak log error jika error bukan dari client
			if websocket.IsUnexpectedCloseError(err, 1001, 1006, 1003) {
				log.Printf("error reading message : %v", err)
			}

			break
		}

		log.Println(typeMsg[messageType])
		log.Println(string(payload))
		for wsclient := range c.server.clients {
			wsclient.egress <- payload
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
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

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Failed to send message %v", err)
			}

			log.Printf("Message sent")
		}

	}

}
