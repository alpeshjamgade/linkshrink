package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"urlshortner/constants"
)

var logger *zap.SugaredLogger
var countTest int32

func CreateLoggerWithCtx(ctx context.Context) *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          constants.LOG_ENCODING,
			EncoderConfig:     encoderCfg,
			OutputPaths:       []string{"stderr"},
			ErrorOutputPaths:  []string{"stderr"},
		}
		logger = zap.Must(config.Build()).Sugar()
		countTest++
		logger.Debugf("created logger, count: %d", countTest)
	}

	if ctx != nil && ctx.Value(constants.TRACE_ID).(string) != "" {
		traceId := ctx.Value(constants.TRACE_ID)
		return logger.WithOptions(zap.Fields(zap.String(string(constants.TRACE_ID), traceId.(string)), zap.String(string(constants.SERVICE), constants.SERVICE)))
	}

	return logger
}

func CreateLogger() *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		config := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          constants.LOG_ENCODING,
			EncoderConfig:     encoderCfg,
			OutputPaths:       []string{"stderr"},
			ErrorOutputPaths:  []string{"stderr"},
		}
		logger = zap.Must(config.Build()).Sugar()
		countTest++
		logger.Debugf("created logger, count: %d", countTest)
	}
	return logger
}

func GetLevel() zap.AtomicLevel {
	switch constants.LOG_LEVEL {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case "dpanic":
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	}
	return zap.NewAtomicLevelAt(zap.InfoLevel)
}
