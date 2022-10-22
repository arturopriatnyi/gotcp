FROM golang:1.19.0-buster

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0
COPY . .

RUN ./bin/golangci-lint run ./tcpserver/... ./tcpclient/...
