// +build js, wasm

package main

import (
	"sync/atomic"
	"syscall/js"
)

var counter uint64

func main() {
	done := make(chan bool, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-done
}

func hit(this js.Value, inputs []js.Value) interface{} {

	atomic.AddUint64(&counter, 1)
	return counter
}

func registerCallbacks() {
	js.Global().Set("hit", js.FuncOf(hit))
}
