package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/marcelofabianov/vita-assist/identity/config"
)

func defineLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func defineEncoderConfig(configFormat string) zapcore.EncoderConfig {
	switch configFormat {
	case "json":
		return zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			TimeKey:     "timestamp",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		}
	default:
		return zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			TimeKey:      "timestamp",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	}
}

func defineOutputConfig(
	cfg config.Log,
	encoderConfig zapcore.EncoderConfig,
	level zap.AtomicLevel,
) (zapcore.Core, error) {
	var core zapcore.Core
	switch cfg.Output {
	case "stdout":
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		writer := zapcore.AddSync(os.Stdout)
		core = zapcore.NewCore(encoder, writer, level)
	case "file":
		file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		writer := zapcore.AddSync(file)
		core = zapcore.NewCore(encoder, writer, level)
	default:
		return nil, fmt.Errorf("invalid log output: %s", cfg.Output)
	}
	return core, nil
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}
