package mw

import (
	"encoding/json"
	"fmt"
	"goblog/utils"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Auth : Authenticate with jwt
// TODO test authentication
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		jwtToken := req.Header.Get("x-auth-token")
		fmt.Println(jwtToken)
		if jwtToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Not authorized",
			})
			return
		}

		secret := os.Getenv("JWT_KEY")

		token, err := jwt.ParseWithClaims(jwtToken, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if claims, ok := token.Claims.(*utils.CustomClaims); ok {
			// fmt.Println("Auth DEBUG Username:", claims.Name, "Expiry: ", claims.ExpiresAt)
			req.Header.Set("user-name", claims.Name)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Not authorized",
				"error":   err.Error(),
			})
			return
		}

		next.ServeHTTP(w, req)
	})
}
