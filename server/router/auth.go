package router

import (
	"os"
	"time"
	
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

func generateJWT(username string) (map[string]string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Minute *2 ).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
    rtClaims := jwt.MapClaims{
		"sub" : username,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	rtoken, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return map[string]string {
		"access_token":tokenString,
		"refresh_token": rtoken,
	}, nil

}


