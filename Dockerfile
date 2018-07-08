FROM golang:alpine
COPY go-mysql-crud .
WORKDIR /go/src/github.com/go-mysql-crud
