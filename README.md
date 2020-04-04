# go-webassembly-tutorial

## WebAssembly Calculator Demo

## Build

```bash
make <experiment>
```

### Serve

```bash
make serve
```

## Editor configuration

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

#### GoLand / Intellij Ultimate

[Configuring GoLand for WebAssembly (Wasm) projects](https://github.com/golang/go/wiki/Configuring-GoLand-for-WebAssembly)
