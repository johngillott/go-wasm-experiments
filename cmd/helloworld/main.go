package main

import "fmt"

// build GOOS=js GOARCH=wasm go build -o main.wasm cmd/helloworld/main.go

func main() {
	fmt.Println("Hello, WebAssembly!")
}
