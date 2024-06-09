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

type LoginServer struct {
	Logger           zerolog.Logger
	WebPort          string
	NatsAddr         string
	MetricsPort      string
	LoginSearchTopic string
	LoginStoreTopic  string
	NatsConn         *nats.Conn
	MysqlAddr        string
	MysqlConn        *sql.DB
}

func main() {
	ls := &LoginServer{
		Logger:           logger.NewLogger(env.LookupStringEnv("LOG_LEVEL", "devbug")),
		WebPort:          env.LookupStringEnv("WEB_PORT", "9001"),
		NatsAddr:         env.LookupStringEnv("NATS_ADDR", "nats.nats.svc.cluster.local:4222"),
		MetricsPort:      env.LookupStringEnv("METRICS_PORT", "9100"),
		LoginSearchTopic: env.LookupStringEnv("LOGIN_SEARCH_TOPIC", "login.search"),
		LoginStoreTopic:  env.LookupStringEnv("LOGIN_STORE_TOPIC", "login.store"),
		MysqlAddr:        env.LookupStringEnv("MYSQL_ADDR", "myqsl.mysql.svc.cluster.local:3306"),
	}

	var err error
	ls.NatsConn, err = nats.Connect(ls.NatsAddr)
	if err != nil {
		ls.Logger.Fatal().Err(err).Msg("could not connect to nats")
	}

	ls.Logger.Info().Msg("connect to mysql")

	ls.MysqlConn, err = sql.Open("mysql", ls.MysqlAddr)
	if err != nil {
		ls.Logger.Fatal().Err(err).Msg("could not connect to mysql")
	}

	ls.Logger.Info().Msg("connect to nats")

	// nats & mysql 종료
	defer ls.NatsConn.Close()
	defer ls.MysqlConn.Close()

	gin.SetMode("release")
	r := gin.New()
	r.GET("/", ls.OnProbe)
	r.POST("/login", ls.OnLogin)
	r.POST("/mining", ls.OnMining)
	r.POST("/user", ls.OnUser)
	r.POST("/join", ls.OnJoin)

	ls.Logger.Info().Msgf("login server start on :%s", ls.WebPort)

	if err := r.Run(fmt.Sprintf(":%s", ls.WebPort)); err != nil {
		ls.Logger.Err(err).Msg("could not run login server")
		return
	}
}
