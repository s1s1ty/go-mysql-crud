#! /bin/sh
set -e

dep ensure -v -update

export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0

go build -v -o dist/go-mysql-crud

docker build -t go-mysql-crud .