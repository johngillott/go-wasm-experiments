# go-webassembly-tutorial

## Hello world Demo

## Build

```bash
GOOS=js GOARCH=wasm go build -o demo/public/main.wasm cmd/helloworld/main.go
```

### Serve

```bash
go run cmd/serve/main.go
```

### Build and serve

```bash
GOOS=js GOARCH=wasm go build -o demo/public/main.wasm cmd/helloworld/main.go && go run cmd/serve/main.go
```
