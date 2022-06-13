package auth

import (
	"errors"
	"net/http"
	"strings"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken is a function that validates a token.
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Token inválido")
	}

	return nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("QUEBROU")
	}

	return []byte(config.SecretKey), nil
}
