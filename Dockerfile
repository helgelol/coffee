FROM golang:1.23.2-bookworm AS api

WORKDIR /app

RUN go install github.com/air-verse/air@latest
CMD [ "air main.go" ]
