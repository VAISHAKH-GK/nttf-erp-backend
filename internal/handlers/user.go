package handlers

import (
	"github.com/MagnaBit/nttf-erp-backend/internal/dto"
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
	var body dto.LoginReq
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	token, err := h.service.Login(body.Username, body.Password)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.LoginRes{Token: token})
}
