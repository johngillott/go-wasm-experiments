// +build js, wasm

package main

import (
	"syscall/js"
	"time"
)

func main() {

	c := make(chan struct{}, 0)

	ws := js.Global().Get("WebSocket").New("ws://localhost:8081/echo")

	ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		ws.Call("send", "Hello, World!\n")

		return nil
	}))

	ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		time.Sleep(5 * time.Second)
		ws.Call("send", "Hello, World!\n")

		return nil
	}))

	<-c
}
