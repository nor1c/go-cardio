package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nor1c/GoJWTMux/configs"
	"github.com/nor1c/GoJWTMux/helpers"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenFromCookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				message := map[string]string{"message": "401 Unauthorized."}
				helpers.ResponseJSON(w, http.StatusUnauthorized, message)
				return
			}
		}

		claims := &configs.JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenFromCookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return configs.JWT_KEY, nil
		})

		if err != nil {
			message := map[string]string{"message": err.Error()}
			helpers.ResponseJSON(w, http.StatusUnauthorized, message)
			return
		}

		if !token.Valid {
			message := map[string]string{"message": "401 Unauthorized"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, message)
			return
		}

		next.ServeHTTP(w, r)
	})
}
