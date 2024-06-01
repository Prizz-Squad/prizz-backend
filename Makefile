include .env
export

build:
	@go build -o bin/api cmd/server.go

run:
	@./bin/api

compile:
	@go build -o bin/report cmd/report_server.go

start:
	@./bin/report

tidy:
	@go mod tidy

migrate:
	@goose -dir db/migrations postgres "$(DB_URL)" up

drop:
	@goose -dir db/migrations postgres "$(DB_URL)" down

up:
	docker-compose up --build

down:
	docker-compose down