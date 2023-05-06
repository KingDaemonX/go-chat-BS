# go-chat-BS

## This is a basic chat server built in golang using websocket

- this server have two basic function

1. Chat interface to communicate with the bot 🤖
2. A Broadcast chatroom for connected users

### How to run application

- clone this repository
- if you have golang installed run :
  🔳 to remove unwanted dependency and install needed dependency

```bash/zsh
go mod tidy
```

- start server in yout terminal

```bash/zsh
go run main.go
```

### Client / User server

##### it is very important to understand how to communicate with this server as a client, so the instruction should follow after starting the server

🔳 This is not some fancy react frontend code as I am a backend developer 😅... so you'd run it in your browser console

- Open Browser
  -Open the browser console using <b>CTRL + SHIFT + I</b>

Input this line of Javascript code:

- First

```
let socket = new WebSocket("ws://localhost:3000/chat")
```

- Next

```
socket.onmessage = event => { console.log("chat server: ", event.data)}
```

With the above code, you're already connected with the server and ready to chat

### To chat Bot :

-

```
socket.send("Hello Bot :)")
```

### To Chat Broadcast / GC

- send regular text using the :

```
socket.send("any text that goes")
```

##### I hopoe you like it, if yes hire or refer me for golang job role
