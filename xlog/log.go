// Package xlog provides extended logging functionality using structured loggers
// with support for multiple log formats, levels, and integration with tracing.
package xlog

import (
	"os"
	"time"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MsgKey is the key used for storing message content in structured logs.
const MsgKey = "msg"

// Init initializes and configures a logger with the specified parameters.
// It sets up consistent logging with metadata like profile, color, service name,
// version, node name, and trace identifiers. Returns the configured logger.
func Init(logType, logLevel string, profile, color, name, version string, nodeName string) (logger log.Logger) {
	var base log.Logger

	switch logType {
	case "zap":
		base = newZapLogger(logLevel)
	default:
		base = log.DefaultLogger
	}

	logger = log.With(base,
		"ts", log.Timestamp(time.DateTime),
		"profile", profile,
		"color", color,
		"caller", log.DefaultCaller,
		"svc", name,
		"sver", version,
		"node", nodeName,
		"trace", tracing.TraceID(),
		"span", tracing.SpanID(),
	)
	log.SetLogger(logger)

	return
}

func newZapLogger(logLevel string) log.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     MsgKey,
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
	), level)
	zlogger := zap.New(core).WithOptions(zap.AddCaller())

	return kzap.NewLogger(zlogger)
}
