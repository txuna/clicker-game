package main

import (
	"clicker/internal/memory"
	"clicker/internal/mq"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func (gc *GameClient) OnLoginStore(msg *nats.Msg) {
	req := &mq.NatsLoginStoreRequest{}
	err := json.Unmarshal(msg.Data, req)
	if err != nil {
		gc.Logger.Err(err).Msg("could not unmarshal login.store msg")
		return
	}

	gc.Logger.Info().Msgf("receive msg userid(%s), token(%s)", req.UserId, req.Token)

	userId := req.UserId + memory.ACCOUNT_PREFIX
	token := req.Token

	if err := memory.Set(gc.RedisConn, userId, token); err != nil {
		gc.Logger.Err(err).Msgf("could not set userid(%s)-token(%s)", userId, token)
	}
}
