package server

import (
	"context"
	"errors"

	"github.com/MagnaBit/nttf-erp-backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

type WebServer struct {
	*fiber.App
	DB *db.Store
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

	var server = &WebServer{
		App: fiber.New(),
		DB:  db,
	}

	return server
}
