package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString([]byte("eniac_is_ishaan"))
}

// func VerifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, hmacKeyfunc)
// 	if err != nil {
// 		return err
// 	}
// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}
// 	return nil
// }

func GetUsernameFromToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("eniac_is_ishaan"), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim missing")
	}

	return username, nil
}