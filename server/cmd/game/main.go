package main

import (
	"clicker/pkg/env"
	"clicker/pkg/logger"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type GameServer struct {
	WebPort          string
	MetricsPort      string
	NatsAddr         string
	MysqlAddr        string
	NatsConn         *nats.Conn
	MysqlConn        *sql.DB
	Logger           zerolog.Logger
	LoginSearchTopic string
}

func main() {
	gs := &GameServer{
		WebPort:          env.LookupStringEnv("WEB_PORT", "9003"),
		NatsAddr:         env.LookupStringEnv("NATS_ADDR", "nats.nats.svc.cluster.local:4222"),
		MysqlAddr:        env.LookupStringEnv("MYSQL_ADDR", "myqsl.mysql.svc.cluster.local:3306"),
		Logger:           logger.NewLogger(env.LookupStringEnv("LOG_LEVEL", "debug")),
		MetricsPort:      env.LookupStringEnv("METRICS_PORT", "9100"),
		LoginSearchTopic: env.LookupStringEnv("LOGIN_SEARCH_TOPIC", "login.search"),
	}

	var err error
	gs.NatsConn, err = nats.Connect(gs.NatsAddr)
	if err != nil {
		gs.Logger.Fatal().Err(err).Msg("could not connect to nats")
	}

	gs.MysqlConn, err = sql.Open("mysql", gs.MysqlAddr)
	if err != nil {
		gs.Logger.Fatal().Err(err).Msg("could not connect to mysql")
	}

	gin.SetMode("release")
	r := gin.New()
	r.GET("/", gs.OnProbe)
	r.POST("/mining", gs.OnMining)
	r.POST("/user", gs.OnUser)
	r.POST("/ranking", gs.OnRanking)

	gs.Logger.Info().Msg("start game server")

	if err := r.Run(fmt.Sprintf(":%s", gs.WebPort)); err != nil {
		gs.Logger.Fatal().Err(err).Msg("could not start run game server")
	}
}
