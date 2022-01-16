package router

import (
	"antiNCP/controller"
	"antiNCP/middleware"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(g *echo.Group) {
	g.GET("/publicKey", controller.UserGetPublicKey)
	g.GET("/token", controller.UserGetToken)
	g.POST("/register", controller.UserRegister)

	g.POST("/status", controller.UserGetStatus, middleware.JWTMiddleware())
}
