package main

import (
	"fmt"
)

func main() {
	// new channel which send 'strings'
	msgChan := make(chan string)
	go func() {
		msgChan <- "Hello"
	}()
	msg := <-msgChan
	fmt.Println(msg)
}
