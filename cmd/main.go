package main

import (
	"os"

	"github.com/MagnaBit/nttf-erp-backend/server"
	"github.com/joho/godotenv"
)

func main() {
	var port string

	godotenv.Load(".env")

	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	server.Run(port)
}
