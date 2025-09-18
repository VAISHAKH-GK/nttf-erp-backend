package server

import (
	"context"
	"errors"
	"os"

	"github.com/MagnaBit/nttf-erp-backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

type WebServer struct {
	*fiber.App
	DB        *db.Store
	JwtSecret string
}

func (s *WebServer) Shutdown(ctx context.Context) error {
	var errs []error
	if err := s.ShutdownWithContext(ctx); err != nil {
		errs = append(errs, err)
	}

	if s.DB != nil {
		s.DB.Close(ctx)
	}

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

	return server
}
