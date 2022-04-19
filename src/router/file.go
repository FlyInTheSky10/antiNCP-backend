package router

import (
	"github.com/FlyInThesky10/antiNCP-backend/controller"
	"github.com/FlyInThesky10/antiNCP-backend/middleware"
	"github.com/labstack/echo/v4"
)

func InitFileRouter(g *echo.Group) {
	g.Use(middleware.JWTMiddleware())
	g.POST("/healthCode", controller.UserUploadHealthCodeFile)
	g.POST("/travelCode", controller.UserUploadTravelCodeFile)
}
