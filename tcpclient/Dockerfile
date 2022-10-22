FROM golang:1.19-alpine

WORKDIR ./tcpclient
COPY . .

RUN go build -o ./build/tcpclient ./cmd/tcpclient/main.go
CMD ["./build/tcpclient"]