# Dev
FROM golang:1.16-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main /usr/src/app/main.go

EXPOSE 9090

ENTRYPOINT ["./main"]
