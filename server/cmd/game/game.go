package main

import (
	"clicker/internal/database"
	"clicker/internal/mq"
	"clicker/pkg/reqres"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"time"

	"github.com/nats-io/nats.go"
)

func (gs *GameServer) VerifyToken(userId, token string) int {
	data, err := json.Marshal(&mq.NatsLoginSearchRequest{
		UserId: userId,
		Token:  token,
	})

	if err != nil {
		gs.Logger.Err(err).Msg("failed marshal NatsLoginSearchRequest")
		return reqres.ERROR_INVALID_REQUEST
	}

	msg, err := gs.NatsConn.RequestMsg(&nats.Msg{
		Subject: gs.LoginSearchTopic,
		Data:    data,
	}, time.Second)

	if err != nil {
		gs.Logger.Err(err).Msgf("failed requestmsg subject:%s", gs.LoginSearchTopic)
		return reqres.ERROR_INTERNAL_SERVER
	}

	res := &mq.NatsLoginSearchResponse{}
	err = json.Unmarshal(msg.Data, res)
	if err != nil {
		gs.Logger.Err(err).Msg("could not unmarshal NatsLoginSearchResponse")
		return reqres.ERROR_INTERNAL_SERVER
	}

	if !res.Result {
		gs.Logger.Err(err).Msgf("none exist userid(%s)", userId)
		return reqres.ERROR_INVALID_USER
	}

	return reqres.ERROR_NONE
}

func (gs *GameServer) DoLoadRanking(req *reqres.RankingRequest) ([]database.Player, int) {
	if ret := gs.VerifyToken(req.UserId, req.Token); ret != reqres.ERROR_NONE {
		gs.Logger.Info().Msgf("failed verify token %s", req.UserId)
		return []database.Player{}, ret
	}

	players, err := database.LoadPlayers(gs.MysqlConn)
	if err != nil {
		gs.Logger.Err(err).Msg("could not load players")
		return players, reqres.ERROR_INTERNAL_SERVER
	}

	return players, reqres.ERROR_NONE
}

func (gs *GameServer) DoLoadUser(req *reqres.UserRequest) (database.Player, int) {
	if ret := gs.VerifyToken(req.UserId, req.Token); ret != reqres.ERROR_NONE {
		gs.Logger.Info().Msgf("failed verify token %s", req.UserId)
		return database.Player{}, ret
	}

	account, err := database.FindAccount(gs.MysqlConn, req.UserId)
	if err != nil {
		gs.Logger.Err(err).Msg("could not find account")
		return database.Player{}, reqres.ERROR_NOT_EXIST_USER
	}

	player, err := database.FindPlayer(gs.MysqlConn, account.Id)
	if err != nil {
		gs.Logger.Err(err).Msg("could not find player")
		return database.Player{}, reqres.ERROR_NOT_EXIST_USER
	}

	return player, reqres.ERROR_NONE
}

func (gs *GameServer) DoMining(req *reqres.MiningRequest) (int, int) {
	if ret := gs.VerifyToken(req.UserId, req.Token); ret != reqres.ERROR_NONE {
		gs.Logger.Info().Msgf("failed verify token %s", req.UserId)
		return 0, ret
	}

	account, err := database.FindAccount(gs.MysqlConn, req.UserId)
	if err != nil {
		gs.Logger.Err(err).Msg("could not find account")
		return 0, reqres.ERROR_NOT_EXIST_USER
	}

	player, err := database.FindPlayer(gs.MysqlConn, account.Id)
	if err != nil {
		gs.Logger.Err(err).Msg("could not find player")
		return 0, reqres.ERROR_NOT_EXIST_USER
	}

	ret, _ := rand.Int(rand.Reader, big.NewInt(10))
	mineCoin := int(ret.Int64() + 1)

	err = database.UpdatePlayerCoin(gs.MysqlConn, player.PlayerId, mineCoin)
	if err != nil {
		gs.Logger.Err(err).Msg("could not update player coin")
		return 0, reqres.ERROR_INTERNAL_SERVER
	}

	return mineCoin, reqres.ERROR_NONE
}
