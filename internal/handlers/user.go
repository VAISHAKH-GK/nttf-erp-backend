package handlers

import (
	"github.com/MagnaBit/nttf-erp-backend/internal/services"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	return c.SendString("LOGIN ROUTE")
}
