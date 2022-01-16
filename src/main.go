package main

import (
	"antiNCP/config"
	. "antiNCP/log"
	middlewareFly "antiNCP/middleware"
	"antiNCP/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Validator = middlewareFly.GetValidator()

	r := e.Group(config.C.App.Root)
	router.InitRouter(r)

	Logger.Fatal(e.Start(config.C.App.Addr))
}
