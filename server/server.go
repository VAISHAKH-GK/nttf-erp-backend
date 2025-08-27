package server

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
)

type WebServer struct {
	*fiber.App
}

func (s *WebServer) Shutdown(ctx context.Context) error {
	var errs []error
	if err := s.ShutdownWithContext(ctx); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func New() *WebServer {
	var server = &WebServer{
		App: fiber.New(),
	}

	return server
}
