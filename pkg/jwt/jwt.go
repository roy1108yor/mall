package jwt

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserId   string
	UserName string

	jwt.RegisteredClaims
}
