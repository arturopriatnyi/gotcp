# Go TCP

Just an exercise of using TCP protocol with Go standard library.

## To test

Start a server with:
```sh
go run ./tcpserver/cmd/main.go 8080
```

Start clients with:
```sh
go run ./tcpclient/cmd/main.go localhost:8080
```
