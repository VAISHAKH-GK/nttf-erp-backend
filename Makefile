# Detect OS
ifeq ($(OS),Windows_NT)
    EXE = bin\api.exe
    RUN = .\$(EXE)
else
    EXE = bin/api
    RUN = ./$(EXE)
endif

sqlc:
	@sqlc generate

build: sqlc
	@go build -o $(EXE) cmd/api/main.go

seed:
	@go run cmd/seed/main.go

run: build
	@$(RUN)

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
