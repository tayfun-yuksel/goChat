package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}


func UpGrade(w http.ResponseWriter, r *http.Request)(*websocket.Conn, error){
	conn, er := upgrader.Upgrade(w, r, nil)
	if er != nil {
		log.Println("error while upgrading connection:   ", er)
		return nil, er 
	}

	return conn, nil
}




