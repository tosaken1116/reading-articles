# Use the official Golang image as the base image
FROM golang:1.22-alpine

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@v1.29.0
WORKDIR /app

CMD ["air", "-c", ".air.toml"]