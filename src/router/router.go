package router

import "github.com/labstack/echo/v4"

func InitRouter(g *echo.Group) {
	rUser := g.Group("/user")
	InitUserRouter(rUser)
	rCode := g.Group("/code")
	InitCodeRouter(rCode)
}
