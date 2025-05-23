package logic

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateJwt(secretKey *rsa.PrivateKey, userId int, expireTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(expireTime).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(secretKey)
}
