package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func jwtSecret() []byte {
	return []byte("eniac_is_ishaan")
}

func hmacKeyfunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return jwtSecret(), nil
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(jwtSecret())
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, hmacKeyfunc)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func VerifyTokenAndGetUsername(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, hmacKeyfunc)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim missing")
	}

	return username, nil
}
