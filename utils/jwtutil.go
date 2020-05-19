package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims : claims object {payload} content
type CustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken : Generates a json web token and returns it as a string
func GenerateToken(username string) (string, error) {
	mySecretKey := []byte(os.Getenv("JWT_KEY"))

	claims := CustomClaims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 360000,
			Issuer:    "GoBlog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		fmt.Println("Error generating token")
		return "", err
	}

	return tokenString, nil

}
