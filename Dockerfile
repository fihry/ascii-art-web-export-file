FROM golang:1.22.3-alpine

RUN apk add bash

COPY  . /app

WORKDIR  /app

LABEL version="0.0.1"
LABEL description="This is a Dockerfile for the ascii-art-web-export-file server."

CMD go run server.go

