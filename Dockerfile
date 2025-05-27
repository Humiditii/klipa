FROM golang:1.24.3-alpine

RUN apk add --no-cache git curl && \
    curl -sSfL https://github.com/goreleaser/goreleaser/releases/latest/download/goreleaser_Linux_x86_64.tar.gz | tar -xz -C /usr/local/bin

WORKDIR /go/src/github.com/Humiditii/klipa
