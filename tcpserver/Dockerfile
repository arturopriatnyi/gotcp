FROM golang:1.19-alpine

WORKDIR ./tcpserver
COPY . .

RUN go build -o ./build/tcpserver ./cmd/tcpserver/main.go
CMD ["./build/tcpserver"]
