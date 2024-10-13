package logger

import (
	"errors"
	"io"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
	LogLevelFatal = "fatal"
)

const (
	logLevelFlag          = "log-level"
	defaultSentryLogLevel = LogLevelError
	defaultLogLevel       = LogLevelDebug
)

func NewSentryFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    logLevelFlag,
			EnvVars: []string{"LOG_LEVEL"},
			Value:   defaultLogLevel,
			Usage:   "Log level for logging",
		},
	}
}

type syncer interface {
	Sync() error
}

// NewFlusher creates a new syncer from given syncer that log a error message if failed to sync.
func NewFlusher(s syncer) func() {
	return func() {
		// ignore the error as the sync function will always fail in Linux
		// https://github.com/uber-go/zap/issues/370
		_ = s.Sync()
	}
}

func ParseLogLevel(logLevel string) (level zapcore.Level, err error) {
	switch logLevel {
	case LogLevelDebug:
		return zapcore.DebugLevel, nil
	case LogLevelInfo:
		return zapcore.InfoLevel, nil
	case LogLevelWarn:
		return zapcore.WarnLevel, nil
	case LogLevelError:
		return zapcore.ErrorLevel, nil
	case LogLevelFatal:
		return zapcore.FatalLevel, nil
	default:
		return level, errors.New("invalid log level")
	}
}

func MustParseLogLevel(logLevel string) zapcore.Level {
	level, err := ParseLogLevel(logLevel)
	if err != nil {
		panic(err)
	}
	return level
}

func newLogger(c *cli.Context) (*zap.Logger, error) {
	var writers = []io.Writer{os.Stdout}
	w := io.MultiWriter(writers...)

	logLevel, err := ParseLogLevel(c.String(logLevelFlag))
	if err != nil {
		return nil, err
	}
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.CallerKey = "caller"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(config)
	cc := zap.New(zapcore.NewCore(encoder, zapcore.AddSync(w), logLevel), zap.AddCaller())

	return cc, nil
}

// NewLogger creates a new sugared logger and a flush function. The flush function should be
// called by consumer before quitting application.
// This function should be used most of the time unless
// the application requires extensive performance.
func NewLogger(c *cli.Context) (logger *zap.Logger, flusher func(), err error) {
	logger, err = newLogger(c)
	if err != nil {
		return nil, nil, err
	}

	return logger, NewFlusher(logger), nil
}
