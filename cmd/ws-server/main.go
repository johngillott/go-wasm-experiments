package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/echo", websocket.Handler(echoServer))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func echoServer(ws *websocket.Conn) {
	defer ws.Close()
	io.Copy(ws, ws)
}
