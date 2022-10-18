# Go TCP

Just an exercise of using TCP protocol with Go standard library.

## To test

Start a server with:
```sh
go run cmd/tcpserver/main.go 8080
```

Start clients with:
```sh
go run cmd/tcpclient/main.go localhost:8080
```
