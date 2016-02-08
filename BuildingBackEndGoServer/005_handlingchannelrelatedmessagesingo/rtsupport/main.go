package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

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
		//		msgType, msg, err := socket.ReadMessage()
		//		if err != nil {
		//			fmt.Println(err)
		//			return
		//		}
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)
		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			go subscribeChannel(socket)
			
		}
		//		fmt.Println(string(msg))
		//		if err = socket.WriteMessage(msgType, msg); err != nil {
		//			fmt.Println(err)
		//			return
		//		}
	}
}

func addChannel(data interface{}) (error) {
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	fmt.Println("added channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn)  {
	// TODO: rethinkDB query / changefeed
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add",
			Channel{"1", "Software Support"},
		}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}