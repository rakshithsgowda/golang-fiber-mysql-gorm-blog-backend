package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

 secretKey := os.Getenv("JWT_SECRET")

func GeneratJWT(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Issuer: issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix()
		})
	return claims.SignedString([]byte(SecretKey))
}
