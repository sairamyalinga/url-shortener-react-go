package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

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
		next.ServeHTTP(w, r)
		
	})

}
