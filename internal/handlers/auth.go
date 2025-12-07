package handlers

import (
	"errors"

	"github.com/Keracode/vidyarthidesk-backend/internal/domain"
	"github.com/Keracode/vidyarthidesk-backend/internal/dto"
	"github.com/Keracode/vidyarthidesk-backend/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	sess := session.FromContext(c)
	var body dto.LoginReq

	userAgent := c.Get("User-Agent")
	ip := c.IP()

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorRes{
			Error:   err.Error(),
			Message: "Invalid  Request Body",
		})
	}

	res, err := h.service.Login(c.Context(), body, userAgent, ip)
	if err != nil {
		h.handleError(c, err)
	}

	sess.Set("refreshToken", res.RefreshToken)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AuthHandler) RefreshToken(c fiber.Ctx) error {
	sess := session.FromContext(c)
	refreshToken, ok := sess.Get("refreshToken").(string)
	if !ok || refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorRes{
			Error:   "Unauthorized",
			Message: "No refresh token found in session",
		})
	}

	res, err := h.service.RefreshToken(c.Context(), refreshToken)
	if err != nil {
		return h.handleError(c, err)
	}

	sess.Set("refreshToken", res.RefreshToken)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AuthHandler) handleError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrInvalidCredentials):
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorRes{
			Error: "Invalid credentials",
		})
	case errors.Is(err, domain.ErrInvalidRefreshToken):
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorRes{
			Error: "Invalid or expired refresh token",
		})
	case errors.Is(err, domain.ErrSessionExpired):
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorRes{
			Error: "Session expired",
		})
	case errors.Is(err, domain.ErrTokenRevoked):
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorRes{
			Error: "Token has been revoked",
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorRes{
			Error: "Internal server error",
		})
	}
}
