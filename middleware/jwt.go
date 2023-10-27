package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kalougata/mall/pkg/e"
	myJwt "github.com/kalougata/mall/pkg/jwt"
	"github.com/kalougata/mall/pkg/response"
	"github.com/spf13/viper"
)

type JWTMiddleware struct {
	conf *viper.Viper
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

func (jm *JWTMiddleware) AdminJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authoritarian")
		if tokenString == "" {
			return response.Build(c, e.ErrUnauthorized(), nil)
		}
		claims, err := myJwt.ParseToken(tokenString, jm.conf.GetString("jwt.admin.secret"))
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				return response.Build(c, e.ErrUnauthorized().WithMsg("token已过期"), nil)
			}
			return response.Build(c, e.ErrUnauthorized().WithMsg("token校验失败"), nil)
		}
		c.Set("userId", claims.ID)
		return c.Next()
	}
}
