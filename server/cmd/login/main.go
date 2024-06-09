package main

import (
	"clicker/pkg/env"
	"clicker/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type LoginServer struct {
	Logger           zerolog.Logger
	WebPort          string
	NatsAddr         string
	MetricsPort      string
	LoginSearchTopic string
	LoginStoreTopic  string
}

func main() {
	ls := &LoginServer{
		Logger:           logger.NewLogger(env.LookupStringEnv("LOG_LEVEL", "devbug")),
		WebPort:          env.LookupStringEnv("WEB_PORT", "9001"),
		NatsAddr:         env.LookupStringEnv("NATS_ADDR", "nats.nats.svc.cluster.local:9003"),
		MetricsPort:      env.LookupStringEnv("METRICS_PORT", "9100"),
		LoginSearchTopic: env.LookupStringEnv("LOGIN_SEARCH_TOPIC", "login.search"),
		LoginStoreTopic:  env.LookupStringEnv("LOGIN_STORE_TOPIC", "login.store"),
	}

	gin.SetMode("release")
	r := gin.New()
	r.GET("/", ls.OnProbe)
	r.POST("/login", ls.OnLogin)
	r.POST("/mining", ls.OnMining)
	r.POST("/user", ls.OnUser)

	ls.Logger.Info().Msgf("login server start on :%s", ls.WebPort)

	if err := r.Run(fmt.Sprintf(":%s", ls.WebPort)); err != nil {
		ls.Logger.Err(err).Msg("could not run login server")
		return
	}
}
