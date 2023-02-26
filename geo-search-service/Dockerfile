# Builder
FROM golang:1.19.5-bullseye as builder

ARG APP_TARGET=api
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Build Go code
FROM builder as build
COPY . .
RUN go build -o out cmd/${APP_TARGET}/main.go

# Runner
FROM debian:stretch-slim

WORKDIR /app

RUN mkdir /app/upload
COPY --from=build /app/out /app/out

ENTRYPOINT ["./out"]
