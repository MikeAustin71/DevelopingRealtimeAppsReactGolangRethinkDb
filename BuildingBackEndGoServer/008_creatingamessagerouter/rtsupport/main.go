package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/*
What does upgrader do ?
Switch Protocols
(http -> websocket)
*/

func main() {
	router := &Router{}

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
