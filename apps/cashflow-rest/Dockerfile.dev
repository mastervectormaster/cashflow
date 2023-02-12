FROM golang:1.18.0-alpine

WORKDIR /root

# Air for hot-reloading
RUN go install github.com/cosmtrek/air@latest

# swag for OpenAPI doc gen
RUN go install github.com/swaggo/swag/cmd/swag@latest

CMD air