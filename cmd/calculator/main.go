package main

import "syscall/js"

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func add(this js.Value, inputs []js.Value) interface{} {

	x := inputs[0].Float()
	y := inputs[1].Float()

	return x + y
}

func subtract(this js.Value, inputs []js.Value) interface{} {
	x := inputs[0].Float()
	y := inputs[1].Float()

	return x - y
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
