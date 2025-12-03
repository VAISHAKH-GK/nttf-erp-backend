package server

import (
	"github.com/MagnaBit/nttf-erp-backend/internal/handlers"
	"github.com/MagnaBit/nttf-erp-backend/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (s *WebServer) RegisterRoutes() {
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	s.App.Get("/", s.HandleIndexRotue)

	api := s.App.Group("/api")

	authService := services.NewAuthService(s.DB.Queries, s.Config.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)

	api.Post("/auth/login", authHandler.Login)
	api.Post("/auth/refresh", authHandler.RefreshToken)
}

func (s *WebServer) HandleIndexRotue(c fiber.Ctx) error {
	return c.SendString("Index Router")
}
