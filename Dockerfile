FROM golang:1.22.0-bookworm as api

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD [ "air main.go" ]