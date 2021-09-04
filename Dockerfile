# syntax=docker/dockerfile:1

# Build

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /api

# Deploy

FROM alpine:3.14

WORKDIR /

COPY --from=build /api /api

COPY --from=build /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

EXPOSE 8080

ENTRYPOINT ["/api"]
