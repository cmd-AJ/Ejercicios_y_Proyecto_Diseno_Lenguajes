package utils

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Sets up the zerolog logger output based on the provided configuration
func ConfigureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level := parseLogLevel("debug")

	zerolog.SetGlobalLevel(level)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "panic":
		return zerolog.PanicLevel
	case "fatal":
		return zerolog.FatalLevel
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		// Default log level
		return zerolog.InfoLevel
	}
}
