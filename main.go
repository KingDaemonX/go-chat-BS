package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type server struct {
	conn map[*websocket.Conn]bool
	msg  []byte
}

func newServer() *server {
	return &server{
		conn: make(map[*websocket.Conn]bool),
		msg:  make([]byte, 1042),
	}
}

func (s *server) startServer(ws *websocket.Conn) {
	fmt.Printf("Starting Chat Server...🌀\n%s just joined the chat room\n", ws.RemoteAddr())

	s.conn[ws] = true

	s.chatRoom(ws)
}

func (s *server) chatRoom(ws *websocket.Conn) {
	for {
		if n, err := ws.Read(s.msg); err != nil {
			if err == io.EOF {
				fmt.Println("[❗]Error : ", err.Error())
				break
			}
			fmt.Println("[❗]Error : ", err.Error())
		} else {
			ws.Write([]byte("--> Hello from KingDaemonX Bot...How may I help you 🤖"))
			fmt.Printf("--> Client Message: %s", s.msg[:n])
		}
	}
}

func main() {
	fmt.Println("Hello world from KingDaemonX 👋 \nThis is a basic chat application backend service built with golang")

	server := newServer()

	http.Handle("/chat", websocket.Handler(server.startServer))

	http.ListenAndServe(":3000", nil)
}
