package logger

import (
	"context"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/configuration"
	domain "github.com/mokmok-dev/golang-template/domain/logger"
	"golang.org/x/exp/slog"
)

var _ domain.Logger = (*Logger)(nil)

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(
	config configuration.Log,
) (*Logger, error) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			switch {
			case attr.Key == slog.MessageKey:
				return slog.String("message", attr.Value.String())
			case attr.Key == slog.LevelKey && attr.Value.String() == slog.LevelWarn.String():
				return slog.String("severity", "WARNING")
			case attr.Key == slog.TimeKey:
				return slog.String("time", attr.Value.Time().Format(time.RFC3339Nano))
			default:
				return attr
			}
		},
	}

	if config.Level == "debug" {
		opts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return &Logger{
		logger: logger,
	}, nil
}

func (l *Logger) Field(key string, iface interface{}) domain.Field {
	return domain.Field{
		Key:       key,
		Interface: iface,
	}
}

func (l *Logger) field(field domain.Field) slog.Attr {
	switch i := field.Interface.(type) {
	case string:
		return slog.String(field.Key, i)
	case int:
		return slog.Int(field.Key, i)
	case bool:
		return slog.Bool(field.Key, i)
	case error:
		return slog.String(field.Key, i.Error())
	default:
		return slog.Any(field.Key, i)
	}
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...domain.Field) {
	slogfields := []any{}
	for _, field := range fields {
		slogfields = append(slogfields, l.field(field))
	}

	l.logger.DebugCtx(ctx, message, slogfields...)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...domain.Field) {
	slogfields := []any{}
	for _, field := range fields {
		slogfields = append(slogfields, l.field(field))
	}

	l.logger.InfoCtx(ctx, message, slogfields...)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...domain.Field) {
	slogfields := []any{}
	for _, field := range fields {
		slogfields = append(slogfields, l.field(field))
	}

	l.logger.ErrorCtx(ctx, message, slogfields...)
}
