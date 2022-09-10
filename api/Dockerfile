# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS development

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
