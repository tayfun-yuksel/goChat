package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tayfun-yuksel/goChat/pkg/websocket"
)



func serve_ws(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println("serving websocket............................... ")
	conn, er := websocket.UpGrade(w, r)
	if er != nil{
		log.Println("error while upgraging serveWs in main.go: ", er)
		return
	}
	client := &websocket.Client{
		Connection: conn,
		Pool: pool,
	}
	pool.Regiser <- client
	client.Read()

}

func setupRoutes()  {
	pool := websocket.NewPool()
	pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serve_ws(pool, w, r)
	})
}

func main()  {
	fmt.Println("goChat v1....")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}




