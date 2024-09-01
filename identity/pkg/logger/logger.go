package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/marcelofabianov/vita-assist/identity/config"
	"github.com/marcelofabianov/vita-assist/identity/internal/core/contract"
)

type Logger struct {
	*zap.Logger
	config *config.Config
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	level := defineLevel(cfg.Log.Level)
	encoderConfig := defineEncoderConfig(cfg.Log.Format)

	core, err := defineOutputConfig(cfg.Log, encoderConfig, level)
	if err != nil {
		return nil, err
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return &Logger{Logger: logger, config: cfg}, nil
}

func (l *Logger) Close() {
	if l.config.Log.Output == "file" {
		if l == nil {
			return
		}
		if err := l.Sync(); err != nil {
			fmt.Printf("error syncing logger: %v\n", err)
		}
	}
}

func (l *Logger) NewMessage(context contract.Context, msg string, fields *[]contract.Field) *contract.Message {
	return &contract.Message{
		Project: l.config.Project,
		Name:    l.config.Name,
		ID:      l.config.ID,
		ENV:     l.config.ENV,
		TZ:      l.config.TZ,
		Context: context,
		Message: msg,
		Fields:  fields,
	}
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

func (l *Logger) Info(context contract.Context, msg string, fields *[]contract.Field) {
	message := l.NewMessage(context, msg, fields)

	l.Logger.Info(message.Message,
		l.String("project", message.Project),
		l.String("name", message.Name),
		l.String("id", message.ID),
		l.String("env", message.ENV),
		l.String("tz", message.TZ),
		l.String("context", message.Context.String()),
		l.Field("fields", message.Fields),
	)
}

func (l *Logger) Error(context contract.Context, msg string, fields *[]contract.Field) {
	message := l.NewMessage(context, msg, fields)

	l.Logger.Error(message.Message,
		l.String("project", message.Project),
		l.String("name", message.Name),
		l.String("id", message.ID),
		l.String("env", message.ENV),
		l.String("tz", message.TZ),
		l.String("context", message.Context.String()),
		l.Field("fields", message.Fields),
	)
}
