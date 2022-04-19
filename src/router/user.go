package router

import (
	"github.com/FlyInThesky10/antiNCP-backend/controller"
	"github.com/FlyInThesky10/antiNCP-backend/middleware"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(g *echo.Group) {
	g.GET("/publicKey", controller.UserGetPublicKey)
	g.GET("/token", controller.UserGetToken)
	g.PUT("/register", controller.UserRegister)

	g.GET("/status", controller.UserGetStatus, middleware.JWTMiddleware())
	g.GET("", controller.UserGetInfo, middleware.JWTMiddleware())
}
