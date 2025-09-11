package main

import (
	"context"
	"encoding/json"
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

func (s *store) Close(ctx context.Context) {
	s.Db.Close(ctx)
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

func loadData(path string, data any) error {
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func addUsers(users []generated.InsertUserParams) {
	for _, user := range users {
		pass, err := utils.HashPassword(user.Password)
		if err != nil {
			log.Printf("Skipping user %s due to hash error: %v", user.Email, err)
			continue
		}

		user.Password = pass

		if err := dbStore.Queries.InsertUser(context.Background(), user); err != nil {
			log.Printf("User insertion failed with error %v", err)
		}
	}

}

func main() {
	var connString string = os.Getenv("GOOSE_DBSTRING")
	connectDb(connString)
	defer dbStore.Close(context.Background())

	var users []generated.InsertUserParams
	if err := loadData("database/seeds/users.json", &users); err != nil {
		log.Fatalf("Error while loading user data %v", err)
	}

	addUsers(users)
}
