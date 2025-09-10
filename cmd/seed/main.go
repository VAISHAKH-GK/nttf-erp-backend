package main

import (
	"context"
	"log"
	"os"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/utils"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

type Store struct {
	Db      *pgx.Conn
	Queries *generated.Queries
}

var DB Store

func connectDb(connString string) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed")
	}

	queries := generated.New(conn)

	DB.Db = conn
	DB.Queries = queries
}

func AddAdminUser() {
	var user generated.InsertUserParams

	pass, err := utils.HashPassword("password")
	if err != nil {
		log.Fatal("Adding admin user failed due to hash failure")
	}

	user.Username = "Admin"
	user.Email = "admin@nttf.co.in"
	user.Password = pass

	if err := DB.Queries.InsertUser(context.Background(), user); err != nil {
		log.Printf("User insertion failed with error %v", err)
		return
	}
}

func main() {
	var connString string = os.Getenv("GOOSE_DBSTRING")
	connectDb(connString)
	defer DB.Db.Close(context.Background())

	AddAdminUser()
}
