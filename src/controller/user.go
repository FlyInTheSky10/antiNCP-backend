package controller

import (
	"antiNCP/controller/param"
	"antiNCP/model"
	"antiNCP/util"
	"antiNCP/util/flyErr"
	"antiNCP/util/response"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UserGetPublicKey(ctx echo.Context) error {
	req := param.ReqUserGetPublicKey{}
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}
	key, found, err := model.GetPrivateKey(req.Id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "fail to get rsa key", err)
	}
	if !found {
		key, err = rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			return response.Error(ctx, http.StatusInternalServerError, "fail to generate rsa key", err)
		}
		err := model.AddPrivateKey(req.Id, key)
		if err != nil {
			return response.Error(ctx, http.StatusInternalServerError, "fail to add rsa key", err)
		}
	}
	publicKey := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKey,
	})
	return response.Success(ctx, param.ResUserGetPublicKey{PublicKey: string(publicKeyPem)})
}

func UserGetToken(ctx echo.Context) error {
	req := param.ReqUserGetToken{}
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}

	user, found, err := model.GetUserByID(req.Id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to find user", err)
	}
	if !found {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find user", err)
	}

	key, found, err := model.GetPrivateKey(req.Id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to find private key", err)
	}
	if !found {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find private key", err)
	}

	pwDecrypt, err := util.RSADecryptFromString(req.Password, key)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to process user info", err)
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), pwDecrypt) != nil {
		return response.Error(ctx, http.StatusBadRequest, "wrong password or user not found", nil)
	}
	token, expire, err := util.GenerateJWTToken(req.Id)
	return response.Success(ctx, param.ResUserGetToken{
		Token:  token,
		Expire: expire,
	})
}
func UserRegister(ctx echo.Context) error {
	req := param.ReqUserRegister{}
	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "bad request", err)
	}

	_, found, err := model.GetUserByID(req.Id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to check user", err)
	}
	if found {
		return response.Error(ctx, http.StatusInternalServerError, "user existed", err)
	}

	key, found, err := model.GetPrivateKey(req.Id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to find private key", err)
	}
	if !found {
		return response.Error(ctx, http.StatusInternalServerError, "cannot find private key", err)
	}

	pwDecrypt, err := util.RSADecryptFromString(req.Password, key)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to process user info", err)
	}

	pwEncrypt, err := bcrypt.GenerateFromPassword(pwDecrypt, bcrypt.DefaultCost)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to process user info", err)
	}

	req.Password = string(pwEncrypt)

	err = model.AddUser(req)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "failed to add user", err)
	}

	return response.Success(ctx, "successfully add user")
}
func UserGetInfo(ctx echo.Context) error {
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

	user.Password = ""
	return response.Success(ctx, user)
}
