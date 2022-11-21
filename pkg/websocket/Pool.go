package websocket

import (
	"fmt"
	"log"
)

type Pool struct{
	Regiser chan *Client
	Unregister chan *Client
	Clients map[*Client]bool
	Broadcast chan *Message
}


func NewPool() *Pool {
	return &Pool{
		Regiser: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan *Message),
	}
}



func (pool *Pool) Start(){
	for{
		select{
		case client := <-pool.Regiser:
			fmt.Printf("user joined ...")
			pool.Clients[client] = true
			for client  := range pool.Clients{
				client.writeMessgeToClient(&Message{Type: 1, Body: "new user joined..." })
			}
			break
		case client := <-pool.Unregister:
			fmt.Println("uesr left")
			pool.Clients[client] = false
			for client := range pool.Clients{
				client.writeMessgeToClient(&Message{Type: 1, Body: "user left...."})
			}
			break
		case message := <- pool.Broadcast:
			fmt.Println("broadcasting mesage:  ", message.Body)
			for client := range pool.Clients{
				client.writeMessgeToClient(message)
			}
		}
	}
}


func (c *Client) writeMessgeToClient(message *Message)  {
	
	if er := c.Connection.WriteJSON(message); er != nil{
		log.Println("error while writing message to client:  ", er)
		return
	}
}
