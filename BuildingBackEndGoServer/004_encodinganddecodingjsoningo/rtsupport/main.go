package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

/*
What does upgrader do ?
Switch Protocols
(http -> websocket)
*/
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	// fmt.Println("Hello from go")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from go")
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// blocks until message received
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
