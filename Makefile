sqlc:
	@sqlc generate

build: sqlc
	@go build -o bin/api cmd/api/main.go

run: build
	@./bin/api

up:
	@goose up

up-by-one:
	@goose up 1

down:
	@goose down

down-by-one:
	@goose down 1

db-reset:
	@goose reset

db-status:
	@goose status
