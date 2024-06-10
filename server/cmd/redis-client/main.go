package main

import (
	"clicker/pkg/env"
	"clicker/pkg/logger"
	"context"

	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type GameClient struct {
	NatsConn        *nats.Conn
	RedisConn       *redis.Client
	LoginStoreTopic string
	Logger          zerolog.Logger
}

func main() {
	gc := &GameClient{
		LoginStoreTopic: env.LookupStringEnv("LOGIN_STORE_TOPIC", "login.store"),
		Logger:          logger.NewLogger(env.LookupStringEnv("LOG_LEVEL", "debug")),
	}

	var err error
	gc.NatsConn, err = nats.Connect(env.LookupStringEnv("NATS_ADDR", "nats.nats.svc.cluster.local:4222"))
	if err != nil {
		gc.Logger.Fatal().Err(err).Msg("could not connect to nats")
	}

	gc.Logger.Info().Msg("connec to nats")

	gc.RedisConn = redis.NewClient(&redis.Options{
		Addr:     env.LookupStringEnv("REDIS_ADDR", "redis-master.redis.svc.cluster.local:6379"),
		Password: "",
		DB:       0,
	})

	// redis connection check
	if err := gc.RedisConn.Ping(context.Background()).Err(); err != nil {
		gc.Logger.Fatal().Err(err).Msg("could not connect to redis")
	}

	gc.Logger.Info().Msg("connect to redis")

	// 토픽 일반적은 pub-sub 구조를 만들면 뒷단의 모든 sub이 같은 메시지를 받음
	// queue sub으로 하나의 consumer만 받을 수 있도록 지정

	_, err = gc.NatsConn.QueueSubscribe(gc.LoginStoreTopic, "login.store", gc.OnLoginStore)
	if err != nil {
		gc.Logger.Fatal().Err(err).Msg("could not subscribe login.store")
	}

	gc.Logger.Info().Msg("start game client redis")

	closeCh := make(chan struct{})
	<-closeCh
	gc.Logger.Info().Msg("close the server")
}
