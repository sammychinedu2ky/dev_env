FROM golang:latest

RUN go install github.com/codegangsta/gin@latest

WORKDIR /app

EXPOSE 3000

CMD gin