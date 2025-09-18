package services

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrTokenGeneration    = errors.New("Token generation failed")
)
