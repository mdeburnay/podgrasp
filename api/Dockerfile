# Dev
FROM golang:latest AS development

WORKDIR /usr/src/app

COPY . .

RUN go build -o main .

EXPOSE 9090

CMD ["/usr/src/app"]
