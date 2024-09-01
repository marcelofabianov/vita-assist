package logger

import (
	"fmt"
	"os"
	"time"

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

type Logger struct {
	*zap.Logger
	config.Log
}

func NewLogger(cfg config.Log) (*Logger, error) {
	level := defineLevel(cfg.Level)
	encoderConfig := defineEncoderConfig(cfg.Format)

	core, err := defineOutputConfig(cfg, encoderConfig, level)
	if err != nil {
		return nil, err
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return &Logger{Logger: logger}, nil
}

func (l *Logger) Close() {
	if l.Log.Output == "file" {
		if l == nil {
			return
		}
		if err := l.Sync(); err != nil {
			fmt.Printf("error syncing logger: %v\n", err)
		}
	}
}

func (l *Logger) Error(err error) zap.Field {
	return zap.Error(err)
}

func (l *Logger) String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func (l *Logger) Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func (l *Logger) Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

func (l *Logger) Field(key string, value any) zap.Field {
	return zap.Any(key, value)
}
