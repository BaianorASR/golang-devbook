package auth

import (
	"time"

	"api/src/config"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken is a function that generates a token.
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()
	permissions["user_id"] = userID
	// secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}
