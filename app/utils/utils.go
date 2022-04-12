package utils

import (
	"fmt"

	"github.com/SantiagoZuluaga/drawflowapi/app/config"
	"github.com/dgrijalva/jwt-go"
)

type Error struct{}

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func (m *Error) Error() string {
	return "Invalid token."
}

func GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	// Create the JWT string
	tokenString, err := token.SignedString(config.JWT_SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(receivedToken string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(receivedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_SECRET, nil
	})

	if err != nil {
		fmt.Println(err)
		return "", &Error{}
	}
	if !token.Valid {
		return "", &Error{}
	}

	return claims.Id, nil
}
