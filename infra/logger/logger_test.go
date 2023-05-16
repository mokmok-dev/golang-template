package logger

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	domain "github.com/mokmok-dev/golang-template/domain/logger"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type fakeClock time.Time

func (c fakeClock) Now() time.Time {
	return time.Time(c)
}

func (c fakeClock) NewTicker(duration time.Duration) *time.Ticker {
	return &time.Ticker{}
}

func newWithWriter(now time.Time, w io.Writer) *zap.Logger {
	clock := fakeClock(now)
	sink := zapcore.AddSync(w)
	lsink := zapcore.Lock(sink)

	enc := zapcore.NewJSONEncoder(encoderConfig())
	core := zapcore.NewCore(enc, lsink, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCallerSkip(1), zap.WithClock(clock))

	return logger
}

func Test_Debug(t *testing.T) {
	setup := func(now time.Time, w io.Writer) *Logger {
		zaplogger := newWithWriter(now, w)

		logger := &Logger{
			logger: otelzap.New(zaplogger),
		}

		return logger
	}

	tests := []struct {
		name    string
		message string
		fields  []domain.Field
		assert  func(time.Time, bytes.Buffer)
	}{
		{
			name:    "ok",
			message: "test message",
			fields: []domain.Field{
				{
					Key:       "key",
					Interface: "value",
				},
			},
			assert: func(now time.Time, buf bytes.Buffer) {
				expect := fmt.Sprintf(`{"severity":"DEBUG","msg":"test message","time":"%s","key":"value"}`, now.Format(time.RFC3339Nano))

				assert.JSONEq(t, expect, buf.String())
			},
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			now := time.Now()
			buf := bytes.Buffer{}

			logger := setup(now, &buf)

			logger.Debug(ctx, test.message, test.fields...)

			test.assert(now, buf)
		})
	}
}

func Test_Info(t *testing.T) {
	setup := func(now time.Time, w io.Writer) *Logger {
		zaplogger := newWithWriter(now, w)

		logger := &Logger{
			logger: otelzap.New(zaplogger),
		}

		return logger
	}

	tests := []struct {
		name    string
		message string
		fields  []domain.Field
		assert  func(time.Time, bytes.Buffer)
	}{
		{
			name:    "ok",
			message: "test message",
			fields: []domain.Field{
				{
					Key:       "key",
					Interface: "value",
				},
			},
			assert: func(now time.Time, buf bytes.Buffer) {
				expect := fmt.Sprintf(`{"severity":"INFO","msg":"test message","time":"%s","key":"value"}`, now.Format(time.RFC3339Nano))

				assert.JSONEq(t, expect, buf.String())
			},
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			now := time.Now()
			buf := bytes.Buffer{}

			logger := setup(now, &buf)

			logger.Info(ctx, test.message, test.fields...)

			test.assert(now, buf)
		})
	}
}

func Test_Error(t *testing.T) {
	setup := func(now time.Time, w io.Writer) *Logger {
		zaplogger := newWithWriter(now, w)

		logger := &Logger{
			logger: otelzap.New(zaplogger),
		}

		return logger
	}

	mockerr := errors.New("mock err")

	tests := []struct {
		name    string
		message string
		fields  []domain.Field
		assert  func(time.Time, bytes.Buffer)
	}{
		{
			name:    "ok",
			message: "test message",
			fields: []domain.Field{
				{
					Key:       "error",
					Interface: mockerr,
				},
			},
			assert: func(now time.Time, buf bytes.Buffer) {
				expect := fmt.Sprintf(`{"severity":"ERROR","msg":"test message","time":"%s","error":"mock err"}`, now.Format(time.RFC3339Nano))

				assert.JSONEq(t, expect, buf.String())
			},
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			now := time.Now()
			buf := bytes.Buffer{}

			logger := setup(now, &buf)

			logger.Error(ctx, test.message, test.fields...)

			test.assert(now, buf)
		})
	}
}
