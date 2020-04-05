// +build js, wasm

package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"

	"golang.org/x/net/websocket"
)

func main() {
	serverAddr := startServer()

	println("WebSocket server listening on ", serverAddr)

	// websocket.Dial()
	client, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	conn, err := websocket.NewClient(newConfig(serverAddr, "/echo"), client)
	if err != nil {
		panic(err)
	}

	msg := []byte("hello, world\n")
	if _, err := conn.Write(msg); err != nil {
		panic(err)
	}
	var gotMsg = make([]byte, 512)
	n, err := conn.Read(gotMsg)
	if err != nil {
		panic(err)
	}
	gotMsg = gotMsg[0:n]
	if !bytes.Equal(msg, gotMsg) {
		panic(err)
	}
	err = conn.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received: %q\n", gotMsg)
}

func startServer() string {
	http.Handle("/echo", websocket.Handler(echoServer))
	server := httptest.NewServer(nil)
	return server.Listener.Addr().String()
}

func echoServer(ws *websocket.Conn) {
	defer ws.Close()
	io.Copy(ws, ws)
}

func newConfig(serverAddr, path string) *websocket.Config {
	config, _ := websocket.NewConfig(fmt.Sprintf("ws://%s%s", serverAddr, path), "http://localhost")
	return config
}
