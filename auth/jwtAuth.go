package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	PhoneNumber int    `json:"phone_number"`
	UserId      uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(phoneNumber int, userId uint64) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		PhoneNumber: phoneNumber,
		UserId:      userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ParseClaims(signedToken string) (int, uint64, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if ok && token.Valid {
		return claims.PhoneNumber, claims.UserId, nil
	} else {
		err = errors.New("token expired")
		return 0, 0, err
	}
}
