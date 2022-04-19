package controller

import (
	"github.com/FlyInThesky10/antiNCP-backend/util"
	"github.com/FlyInThesky10/antiNCP-backend/util/flyErr"
	"github.com/FlyInThesky10/antiNCP-backend/util/response"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func UserUploadHealthCodeFile(ctx echo.Context) error {
	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	file, err := ctx.FormFile("hc")
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find file", err)
	}

	src, err := file.Open()
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot open file", err)
	}
	defer src.Close()

	dst, err := os.Create("code/hc/" + file.Filename)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot upload file", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot upload file", err)
	}

	return response.Success(ctx, "successfully upload file")
}
func UserUploadTravelCodeFile(ctx echo.Context) error {
	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	file, err := ctx.FormFile("tc")
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find file", err)
	}

	src, err := file.Open()
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot open file", err)
	}
	defer src.Close()

	dst, err := os.Create("code/tc/" + file.Filename)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot upload file", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "cannot upload file", err)
	}

	return response.Success(ctx, "successfully upload file")
}
