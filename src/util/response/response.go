package response

import (
	"github.com/FlyInThesky10/antiNCP-backend/config"
	. "github.com/FlyInThesky10/antiNCP-backend/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Info    string      `json:"info"`
}

func Success(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, Response{
		Data:    data,
		Success: true,
		Info:    "",
	})
}

func Error(ctx echo.Context, status int, info string, err error) error {
	if err != nil && config.C.Debug {
		Logger.Panic(err)
	}
	return ctx.JSON(status, Response{
		Data:    nil,
		Success: false,
		Info:    info,
	})
}
