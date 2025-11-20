package handlers

import (
	"github.com/MagnaBit/nttf-erp-backend/internal/dto"
	"github.com/MagnaBit/nttf-erp-backend/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	sess := session.FromContext(c)
	var body dto.LoginReq

	userAgent := c.Get("User-Agent")
	ip := c.IP()

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request Body"})
	}

	authToken, refreshToken, err := h.service.Login(body, userAgent, ip)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	sess.Set("refreshToken", refreshToken)
	return c.Status(fiber.StatusOK).JSON(dto.LoginRes{AuthToken: authToken, RefreshToken: refreshToken})
}

func (h *UserHandler) RefreshToken(c fiber.Ctx) error {
	sess := session.FromContext(c)
	refreshToken := sess.Get("refreshToken").(string)

	authToken, newRefreshToken, err := h.service.RefreshToken(refreshToken)
	if err != nil {
		if err == services.ErrInvalidRefreshToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "session expired"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}

	sess.Set("refreshToken", newRefreshToken)
	return c.Status(fiber.StatusOK).JSON(dto.LoginRes{AuthToken: authToken, RefreshToken: newRefreshToken})
}
