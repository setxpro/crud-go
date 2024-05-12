package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

// Logger config
func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()}, // stdout or file
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()

}

// Get output logger
func getOutputLogs() string {
	outPut := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))

	if outPut == "" {
		return "stdout"
	}
	return outPut
}

// Config Level
func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}
