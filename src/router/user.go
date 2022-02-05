package router

import (
	"antiNCP/controller"
	"antiNCP/middleware"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(g *echo.Group) {
	g.GET("/publicKey", controller.UserGetPublicKey)
	g.GET("/token", controller.UserGetToken)
	g.PUT("/register", controller.UserRegister)

	g.GET("/status", controller.UserGetStatus, middleware.JWTMiddleware())
	g.GET("", controller.UserGetInfo, middleware.JWTMiddleware())
}
