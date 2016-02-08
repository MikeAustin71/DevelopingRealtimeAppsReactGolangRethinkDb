package main
import (
	"net/http"
	"github.com/gorilla/websocket"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct {
	rules map[string]Handler
}


func NewRouter() *Router  {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(msgName string, handler Handler)  {
	r.rules[msgName] = handler
}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	socket, err := upgrader.Upgrade(w, r, nil)

}

