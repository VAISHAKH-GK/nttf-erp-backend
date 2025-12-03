package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/MagnaBit/nttf-erp-backend/config"
	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/MagnaBit/nttf-erp-backend/pkg/hash"
	"github.com/google/uuid"
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

func addUsers(users []generated.InsertUserParams) []uuid.UUID {
	var ids []uuid.UUID

	for _, user := range users {
		pass, err := hash.HashPassword(user.Password)
		if err != nil {
			log.Printf("Skipping user %s due to hash error: %v", user.Email, err)
			continue
		}

		user.Password = pass

		id, err := dbStore.Queries.InsertUser(context.Background(), user)
		if err != nil {
			log.Printf("User insertion failed with error %v", err)
			continue
		}

		ids = append(ids, id)
	}

	return ids
}

func addAccountTypes(accountTypes []generated.InsertAccountTypeParams) map[string]uuid.UUID {
	var ids = make(map[string]uuid.UUID)

	for _, accountType := range accountTypes {
		id, err := dbStore.Queries.InsertAccountType(context.Background(), accountType)
		if err != nil {
			log.Printf("Account Type insertion failed with error %v", err)
			continue
		}

		ids[accountType.Name] = id
	}

	return ids
}

func addUserAccountTypes(userAccountTypes []generated.InsertUserAccountTypeParams) {
	for _, userAccountType := range userAccountTypes {
		if err := dbStore.Queries.InsertUserAccountType(context.Background(), userAccountType); err != nil {
			log.Printf("Account Type insertion failed with error %v", err)
		}
	}
}

func seedUserAndAccounts() {
	var users []generated.InsertUserParams
	if err := loadData("database/seeds/users.json", &users); err != nil {
		log.Fatalf("Error while loading user data %v", err)
	}

	var accountTypes []generated.InsertAccountTypeParams
	if err := loadData("database/seeds/account_types.json", &accountTypes); err != nil {
		log.Fatalf("Error while loading account types%v", err)
	}

	userIds := addUsers(users)
	accountTypeIds := addAccountTypes(accountTypes)

	var userAccountTypes []generated.InsertUserAccountTypeParams

	if len(userIds) > 0 && accountTypeIds["admin"] != uuid.Nil {
		userAccountTypes = append(userAccountTypes, generated.InsertUserAccountTypeParams{UserID: userIds[0], AccountTypeID: accountTypeIds["admin"]})

		userAccountTypes = append(userAccountTypes, generated.InsertUserAccountTypeParams{UserID: userIds[0], AccountTypeID: accountTypeIds["teacher"]})

		userAccountTypes = append(userAccountTypes, generated.InsertUserAccountTypeParams{UserID: userIds[1], AccountTypeID: accountTypeIds["teacher"]})

		userAccountTypes = append(userAccountTypes, generated.InsertUserAccountTypeParams{UserID: userIds[2], AccountTypeID: accountTypeIds["student"]})

		addUserAccountTypes(userAccountTypes)
	}
}

func main() {
	cfg := config.Load()
	connectDb(cfg.DBString)
	defer dbStore.Close(context.Background())

	seedUserAndAccounts()
}
