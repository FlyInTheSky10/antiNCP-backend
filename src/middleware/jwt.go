package middleware

import (
	"antiNCP/config"
	"antiNCP/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &util.JwtClaims{},
		SigningKey: []byte(config.C.Jwt.Secret),
	})
}
