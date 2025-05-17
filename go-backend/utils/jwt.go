package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJWT(email string) (string, error) {
	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		return "", fiber.NewError(fiber.StatusInternalServerError, "JWT_KEY no definido en el entorno")
	}

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
