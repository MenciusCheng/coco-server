package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var DefaultLogger = InitDefaultLogger()

func InitDefaultLogger() *zap.Logger {
	//logger := zap.NewExample()
	//zap.NewProduction()

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "logger",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	logger := zap.New(core)
	//logger = logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	return logger
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	DefaultLogger.Fatal(msg, fields...)
}

func Sync() error {
	return DefaultLogger.Sync()
}

func Logger() *zap.Logger {
	return DefaultLogger
}
