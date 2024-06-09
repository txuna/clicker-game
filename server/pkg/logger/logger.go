package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(logLevel string) zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	_ = logLevel
	return logger
}
