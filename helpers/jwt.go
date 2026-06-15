package helpers

import (
	"os"
	"theatre-test-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID string, user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"user":    user,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
