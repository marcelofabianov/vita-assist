package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/marcelofabianov/vita-assist/identity/config"
)

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
