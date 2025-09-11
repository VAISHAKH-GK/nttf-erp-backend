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

type store struct {
	Db      *pgx.Conn
	Queries *generated.Queries
}

var users []generated.InsertUserParams = []generated.InsertUserParams{
	{
		Username: "Admin",
		Email:    "admin@nttf.co.in",
		Password: "password",
	},
}

var dbStore store

func connectDb(connString string) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed")
	}

	queries := generated.New(conn)

	dbStore.Db = conn
	dbStore.Queries = queries
}

func addUsers() {
	for _, user := range users {
		pass, err := utils.HashPassword(user.Password)
		if err != nil {
			log.Fatal("Adding admin user failed due to hash failure")
		}

		user.Password = pass

		if err := dbStore.Queries.InsertUser(context.Background(), user); err != nil {
			log.Printf("User insertion failed with error %v", err)
			return
		}
	}

}

func main() {
	var connString string = os.Getenv("GOOSE_DBSTRING")
	connectDb(connString)
	defer dbStore.Db.Close(context.Background())

	addUsers()
}
