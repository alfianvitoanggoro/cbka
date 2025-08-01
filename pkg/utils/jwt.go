package utils

import (
	"go-kafka/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, email, role string) (string, error) {
	secretKey := []byte(config.GetAppJWTKey())

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
