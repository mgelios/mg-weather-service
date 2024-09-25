FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . .

RUN go mod download
RUN go build -o ./mg-weather-service

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/mg-weather-service ./mg-weather-service
CMD ("/app/mg-weather-service") 