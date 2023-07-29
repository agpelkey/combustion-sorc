# syntax=docker/dockerfile:1
# Service for inventory management 

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /inventory


CMD ["/inventory"]


