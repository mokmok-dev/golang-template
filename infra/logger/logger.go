package logger

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/mokmok-dev/golang-template/domain/configuration"
	domain "github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ domain.Logger = (*Logger)(nil)

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

type Logger struct {
	logger *otelzap.Logger
}

func NewLogger(
	config configuration.Log,
) (*Logger, error) {
	zapconfig := zap.NewProductionConfig()
	zapconfig.EncoderConfig = encoderConfig()
	switch config.Level {
	case "debug":
		zapconfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	l, err := zapconfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return &Logger{
		logger: otelzap.New(l),
	}, nil
}

func encoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.LevelKey = "severity"
	cfg.EncodeLevel = EncodeLevel
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func (l *Logger) Field(key string, iface interface{}) domain.Field {
	return domain.Field{
		Key:       key,
		Interface: iface,
	}
}

func (l *Logger) field(field domain.Field) zap.Field {
	switch i := field.Interface.(type) {
	case error:
		return zap.Error(i)
	case string:
		return zap.String(field.Key, i)
	case int:
		return zap.Int(field.Key, i)
	case bool:
		return zap.Bool(field.Key, i)
	default:
		return zap.Any(field.Key, i)
	}
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...domain.Field) {
	zapfields := []zap.Field{}
	for _, field := range fields {
		zapfields = append(zapfields, l.field(field))
	}

	l.logger.Ctx(ctx).Debug(message, zapfields...)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...domain.Field) {
	zapfields := []zap.Field{}
	for _, field := range fields {
		zapfields = append(zapfields, l.field(field))
	}

	l.logger.Ctx(ctx).Info(message, zapfields...)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...domain.Field) {
	zapfields := []zap.Field{}
	for _, field := range fields {
		zapfields = append(zapfields, l.field(field))
	}

	l.logger.Ctx(ctx).Error(message, zapfields...)
}
