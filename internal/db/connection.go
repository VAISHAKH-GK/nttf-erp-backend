package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Keracode/vidyarthidesk-backend/internal/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Pool    *pgxpool.Pool
	Queries *generated.Queries
}

func ConnectDB(connString string, maxConn int) *Store {
	config, _ := pgxpool.ParseConfig(connString)
	config.MaxConns = int32(maxConn)

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create DB pool %v", err)
		os.Exit(1)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Database Connection failed \n%v", err)
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
