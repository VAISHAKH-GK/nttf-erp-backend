package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(secret string, claim map[string]any) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claim))

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil
}
