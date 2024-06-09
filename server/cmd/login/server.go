package main

import (
	"clicker/pkg/reqres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ls *LoginServer) OnProbe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (ls *LoginServer) OnLogin(ctx *gin.Context) {
	req := &reqres.LoginRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ls.Logger.Err(err).Msg("invalid login request")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INVALID_REQUEST,
		})
		return
	}

	// 로그인 유저 검사 -> Nats request replay

	// userid - token 저장 요청 -> Nats pub sub

	// 로그인 성공
	token, err := uuid.NewV7()
	if err != nil {
		ls.Logger.Err(err).Msg("could not generate token")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_FAILED_GENERATE_TOKEN,
		})
		return
	}

	ls.Logger.Err(err).Msgf("Login user id{%s} - token{%s}", req.UserId, token.String())
	ctx.JSON(http.StatusOK, gin.H{
		"error_code": reqres.ERROR_NONE,
		"token":      token.String(),
	})
}

func (ls *LoginServer) OnMining(ctx *gin.Context) {

}

func (ls *LoginServer) OnUser(ctx *gin.Context) {

}
