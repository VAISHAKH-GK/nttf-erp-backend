package server

import (
	"context"
	"errors"
	"os"
	"sync"

	"github.com/MagnaBit/nttf-erp-backend/internal/db"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type WebServer struct {
	*fiber.App
	DB        *db.Store
	JwtSecret string
}

func (s *WebServer) Shutdown(ctx context.Context) error {
	var wg sync.WaitGroup
	var errs []error

	wg.Go(func() {
		if err := s.ShutdownWithContext(ctx); err != nil {
			errs = append(errs, err)
		}
	})

	if s.DB != nil {
		wg.Go(func() {
			s.DB.Close(ctx)
		})
	}

	wg.Wait()

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func New() *WebServer {
	db := db.ConnectDB()
	secret := os.Getenv("JWT_SECRET")

	var server = &WebServer{
		App:       fiber.New(),
		DB:        db,
		JwtSecret: secret,
	}

	server.Use(session.New())

	return server
}
