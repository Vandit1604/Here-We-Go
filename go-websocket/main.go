package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is normal http route")
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Socket Updation failed", err)
	}
	defer conn.Close()

	log.Println("Client connected")

	err = conn.WriteMessage(1, []byte("Hi Client"))
	if err != nil {
		log.Fatal("Unable to send message to client")
	}

	for {
		num, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Unable to read message", err)
		}

		// echo back the message
		err = conn.WriteMessage(1, []byte(msg))

		log.Println(num, string(msg))
	}
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/ws", WebsocketHandler)
	http.ListenAndServe(":8080", nil)
}
