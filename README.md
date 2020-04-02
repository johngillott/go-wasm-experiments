# go-webassembly-tutorial

## WebAssembly Calculator Demo

## Build

make build

### Serve

make serve

### Build and serve

make build serve

## Editor Config

### `could not import syscall/js (no package for import syscall/js)`

#### VS Code

When using Go modules related to WebAssembly, namely syscall/js, the default settings in VS Code will trigger error reports like "Build constraints exclude all Go files" in the editor.

Preferences > Settings > Workspace > Go tools env var.

Set the following:

```json
{
    "go.toolsEnvVars": {
        "GOOS": "js",
        "GOARCH": "wasm"
    }
}
```

Ctrl + Shift + P > Reload Window.
