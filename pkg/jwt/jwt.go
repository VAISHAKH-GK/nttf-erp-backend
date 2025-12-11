package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserId   uuid.UUID
	Name     string
	Email    string
	Expiry   time.Time
	IssuedAt time.Time
}

func GenerateJwtToken(secret string, claim Claims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    claim.UserId.String(),
		"name":  claim.Name,
		"email": claim.Email,
		"iat":   claim.IssuedAt.Unix(),
		"exp":   claim.Expiry.Unix(),
	})

	return t.SignedString([]byte(secret))
}
