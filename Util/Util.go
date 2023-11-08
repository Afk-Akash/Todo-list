package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(issuer string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret_key"))

}

func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims.Issuer, nil
}
