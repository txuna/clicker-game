package main

import (
	"clicker/internal/mq"
	"clicker/pkg/reqres"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func (ls *LoginServer) OnProbe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (ls *LoginServer) OnJoin(ctx *gin.Context) {
	req := &reqres.JoinRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ls.Logger.Err(err).Msg("invalid join request")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INVALID_REQUEST,
		})

		return
	}

	code := ls.DoJoin(req)
	if code != reqres.ERROR_NONE {
		ls.Logger.Info().Msgf("failed join: %d", code)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": code,
		})

		return
	}

	ls.Logger.Info().Msgf("Join user(%s)", req.UserId)

	ctx.JSON(http.StatusOK, gin.H{
		"error_code": code,
	})
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

	token, code := ls.DoLogin(req)
	if code != reqres.ERROR_NONE {
		ls.Logger.Info().Msgf("failed login: %d", code)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": code,
		})

		return
	}

	data, err := json.Marshal(&mq.NatsLoginStoreRequest{
		UserId: req.UserId,
		Token:  token,
	})

	if err != nil {
		ls.Logger.Err(err).Msg("could not marshal login store request")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INTERNAL_SERVER,
		})

		return
	}

	err = ls.NatsConn.PublishMsg(&nats.Msg{
		Subject: ls.LoginStoreTopic,
		Data:    data,
	})

	if err != nil {
		ls.Logger.Err(err).Msg("could not publish login store")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INTERNAL_SERVER,
		})

		return
	}

	ls.Logger.Info().Msgf("Login user id{%s} - token{%s}", req.UserId, token)
	ctx.JSON(http.StatusOK, gin.H{
		"error_code": reqres.ERROR_NONE,
		"user_id":    req.UserId,
		"token":      token,
	})
}
