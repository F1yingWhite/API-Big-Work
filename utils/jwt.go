package utils

import (
	"API_BIG_WORK/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	ID string `json:"ID"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 2

func CreateToken(id string) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), JWT{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		},
	}).SignedString([]byte(config.CFG.JWTSigningString))
}

// 解码token
func ParseToken(tokenString string) (*JWT, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.CFG.JWTSigningString), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWT)
	if !ok {
		return nil, err
	}
	return claims, nil
}
