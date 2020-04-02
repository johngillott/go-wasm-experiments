.PHONY: init
init:
		cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./public/

.PHONY: build
build:
		GOOS=js GOARCH=wasm go build -o public/main.wasm cmd/helloworld/main.go

.PHONY: serve
serve:
		go run cmd/serve/main.go

.PHONY: open
open: 
		gio open https://localhost:8080