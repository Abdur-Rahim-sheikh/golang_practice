package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub         int    `json:"sub"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
	jwt.RegisteredClaims
}

func CreateJwt(secret string, data Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, data)
	return token.SignedString([]byte(secret))
}

func VerifyJwt(secret string, tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Explicitly enforce the signing method — never trust the token's own header for this
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
