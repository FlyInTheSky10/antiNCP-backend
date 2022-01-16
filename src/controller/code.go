package controller

import (
	"antiNCP/controller/param"
	"antiNCP/model"
	"antiNCP/util"
	"antiNCP/util/flyErr"
	"antiNCP/util/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserSubmitCode(ctx echo.Context) error {
	req := param.ReqUserSubmitCode{}
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}

	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	err := model.AddSubmission(req, id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to add submission", err)
	}

	return response.Success(ctx, "successfully submitted")
}
func UserViewSubmission(ctx echo.Context) error {
	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	submissions := make([]param.Submission, 0)
	submissions, err := model.GetSubmissionById(id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to view submission", err)
	}

	return response.Success(ctx, submissions)
}
func UserGetStatus(ctx echo.Context) error {
	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	code, err := model.GetStatus(id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to get status", err)
	}

	return response.Success(ctx, code)
}
func AdminGetSubmission(ctx echo.Context) error {
	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	user, found, err := model.GetUserByID(id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to find user", err)
	}
	if !found {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find user", err)
	}

	if user.Admin == 0 {
		return response.Error(ctx, http.StatusForbidden, "permission denied", err)
	}

	submissions, err := model.GetSubmissions()
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to get submissions", err)
	}
	return response.Success(ctx, submissions)
}
func AdminVerifySubmission(ctx echo.Context) error {
	req := param.ReqAdminVerifyCode{}
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}

	id := util.GetIdFromJWT(ctx)
	if id == "" {
		return flyErr.Error{Text: "invalid token"}
	}

	user, found, err := model.GetUserByID(id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to find user", err)
	}
	if !found {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find user", err)
	}

	if user.Admin == 0 {
		return response.Error(ctx, http.StatusForbidden, "permission denied", err)
	}

	err = model.UpdateStatus(req.I, id, req.Status)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to update status", err)
	}

	return response.Success(ctx, "successfully verified")
}
