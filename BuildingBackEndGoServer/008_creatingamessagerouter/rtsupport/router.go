package main
import (
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}


type Router struct {}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	socket, err := upgrader.Upgrade(w, r, nil)

}

