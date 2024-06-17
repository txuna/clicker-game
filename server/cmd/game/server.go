package main

import (
	"clicker/pkg/reqres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gs *GameServer) OnProbe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (gs *GameServer) OnMining(ctx *gin.Context) {
	req := &reqres.MiningRequest{}
	if err := ctx.BindJSON(req); err != nil {
		gs.Logger.Err(err).Msg("invalid join request")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INVALID_REQUEST,
		})

		return
	}

	coin, ret := gs.DoMining(req)

	if ret != reqres.ERROR_NONE {
		gs.Logger.Info().Msgf("user(%s) failed mining", req.UserId)
	} else {
		gs.Logger.Info().Msgf("user(%s) mining coin: %d", req.UserId, coin)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error_code": ret,
		"coin":       coin,
	})
}

func (gs *GameServer) OnUser(ctx *gin.Context) {
	req := &reqres.UserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		gs.Logger.Err(err).Msg("invalid join request")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INVALID_REQUEST,
		})

		return
	}

	player, ret := gs.DoLoadUser(req)
	if ret != reqres.ERROR_NONE {
		gs.Logger.Info().Msgf("failed load user: %d", ret)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": ret,
		})

		return
	}

	gs.Logger.Info().Msgf("load user: %v", player)

	ctx.JSON(http.StatusOK, gin.H{
		"error_code": ret,
		"coin":       player.Coin,
		"max_coin":   player.MaxCoin,
	})
}

func (gs *GameServer) OnRanking(ctx *gin.Context) {
	req := &reqres.RankingRequest{}
	if err := ctx.BindJSON(req); err != nil {
		gs.Logger.Err(err).Msg("invalid join request")
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": reqres.ERROR_INVALID_REQUEST,
		})

		return
	}

	players, ret := gs.DoLoadRanking(req)
	if ret != reqres.ERROR_NONE {
		gs.Logger.Info().Msgf("failed load ranking: %d", ret)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": ret,
		})

		return
	}

	gs.Logger.Info().Msg("load rankings")
	ctx.JSON(http.StatusOK, gin.H{
		"error_code": ret,
		"players":    players,
	})
}
