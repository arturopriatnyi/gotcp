run:
	go run cmd/tcpserver/main.go 10000 &
	go run cmd/tcpclient/main.go localhost:10000 &
	go run cmd/tcpclient/main.go localhost:10000
