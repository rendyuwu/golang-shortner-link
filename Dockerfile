# syntax=docker/dockerfile:1

FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /shortner-api

EXPOSE 80

CMD [ "/shortner-api" ]