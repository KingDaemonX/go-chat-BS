package main

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type server struct {
	conn map[*websocket.Conn]bool
}

func NewServer() server {
	return *&server{
		conn: make(map[*websocket.Conn]bool),
	}
}

func main() {
	fmt.Println("Hello world from KingDaemonX 👋 \n This is a basic chat application backend service built with golang")
}
