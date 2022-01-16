package util

import (
	"antiNCP/config"
	"antiNCP/constant"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

type JwtClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWTToken(id string) (string, int64, error) {
	claims := &JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.JwtExpiresDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.C.Jwt.Secret))
	if err != nil {
		return t, 0, err
	}
	return t, claims.ExpiresAt, nil
}
func GetIdFromJWT(ctx echo.Context) string {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	return claims.Id
}
