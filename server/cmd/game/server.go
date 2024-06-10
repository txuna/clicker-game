package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gs *GameServer) OnProbe(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (gs *GameServer) OnMining(ctx *gin.Context) {

}
