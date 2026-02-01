package logging

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

var (
	logger logr.Logger
)

// Setups logger instance
func Logger() logr.Logger {
	if logger.GetSink() != nil {
		return logger
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zerologr.SetMaxV(1)

	zl := zerolog.New(os.Stderr)
	zl = zl.With().Caller().Timestamp().Logger()
	logger = zerologr.New(&zl)
	return logger
}
