FROM golang:latest

WORKDIR /api

COPY . .

RUN go mod tidy

RUN go build -o bin/api cmd/report_server.go

EXPOSE 5555

CMD ./bin/report