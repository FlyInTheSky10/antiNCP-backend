package router

import (
	"antiNCP/controller"
	"antiNCP/middleware"
	"github.com/labstack/echo/v4"
)

func InitFileRouter(g *echo.Group) {
	g.Use(middleware.JWTMiddleware())
	g.POST("/healthCode", controller.UserUploadHealthCodeFile)
	g.POST("/travelCode", controller.UserUploadTravelCodeFile)
}
