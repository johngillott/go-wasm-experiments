.PHONY: serve
serve:
		@go run cmd/serve/main.go

clean:
		@rm -f ./public/*

.PHONY: calculator
calculator: clean
		@cp ./calculator/index.html ./public/index.html
		@GOOS=js GOARCH=wasm go build -o public/main.wasm calculator/main.go
		@cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./public/wasm_exec.js

.PHONY: channels
channels: clean
		@GOOS=js GOARCH=wasm go build -o ./public/test.wasm ./channels/main.go
		@cp $$(go env GOROOT)/misc/wasm/wasm_exec.html ./public/index.html
		@cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./public/wasm_exec.js

