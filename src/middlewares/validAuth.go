package middlewares

import (
	"api/src/auth"
	"api/src/response"
	"net/http"
)

// ValidAuth is a middleware that valiates user's token.
func ValidAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
