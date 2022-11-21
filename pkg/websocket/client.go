package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id string
	Connection *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	 
	defer func ()  {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	messageType, p, er := c.Connection.ReadMessage()
	if er != nil{
		log.Println("error while reading message from client", er)
	}

	message := &Message{Type: messageType, Body: string(p)}
	c.Pool.Broadcast <- message
	fmt.Println("received message from client----  ")
}
