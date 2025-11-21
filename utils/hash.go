package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHash(password string, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
