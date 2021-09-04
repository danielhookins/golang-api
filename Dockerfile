# syntax=docker/dockerfile:1

# Build

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o api

# Deploy

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]
