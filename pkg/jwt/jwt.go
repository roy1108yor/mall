package jwt

import (
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId   uint
	UserName string

	jwt.RegisteredClaims
}

func GenToken(claims CustomClaims, expiresAt time.Time, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserId:   claims.UserId,
		UserName: claims.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			Audience:  []string{},
		},
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string, secret string) (*CustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
