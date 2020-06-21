package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Print("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}

	go func() {
		defer func() {
			conn.Close()
		}()
		for i := 0; i < 5; i++ {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", i)))
			time.Sleep(500 * time.Millisecond)
		}
	}()
}
