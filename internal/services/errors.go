package services

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrTokenGeneration     = errors.New("token generation failed")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
	ErrDatabase            = errors.New("database request failed")
)
