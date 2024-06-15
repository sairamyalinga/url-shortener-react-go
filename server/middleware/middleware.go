package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))
type ContextKeyType string
const UsernameContextKey ContextKeyType = "username"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
            return
        }
		authHeader = strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        username, ok := claims["sub"].(string)

		if !ok {
			http.Error(w, "Invalid username in token", http.StatusUnauthorized)
			return
		}
		
		ctx := context.WithValue(r.Context(), UsernameContextKey, username)

		next.ServeHTTP(w, r.WithContext(ctx))
		
	})

}
