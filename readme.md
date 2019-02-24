```
GOARCH=wasm GOOS=js go build -o test.wasm main.go
go run server.go
```

[http://localhost:8080/wasm_exec.html](http://localhost:8080/wasm_exec.html)
