FROM golang:1.21.0-alpine

LABEL maintainer="Gabriel Quinteros <gquinteros@sort-boks.com>"

WORKDIR /app

COPY main.go .
COPY .env .
COPY go.mod .
COPY go.sum .
COPY vars ./vars
COPY routes ./routes
COPY controller ./controller

RUN go build -o demomovimientos

EXPOSE 8002

CMD ["./demomovimientos"]