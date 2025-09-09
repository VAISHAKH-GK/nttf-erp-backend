package db

import (
	"context"
	"fmt"
	"os"

	"github.com/MagnaBit/nttf-erp-backend/internal/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Pool    *pgxpool.Pool
	Queries *generated.Queries
}

func ConnectDB() *Store {
	config, _ := pgxpool.ParseConfig(os.Getenv("GOOSE_DBSTRING"))
	config.MaxConns = 5

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Database Connection failed %v", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	return &Store{
		Pool:    pool,
		Queries: generated.New(pool),
	}
}

func (s *Store) Close(ctx context.Context) {
	s.Pool.Close()
}
