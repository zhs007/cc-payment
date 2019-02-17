package logger

import (
	"encoding/json"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/zhs007/cc-payment/config"
	"github.com/zhs007/cc-payment/err"
)

func getLogLevel(str string) zapcore.Level {
	switch str {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.ErrorLevel
	}
}

var logger *zap.Logger
var onceLogger sync.Once

func initLogger() (*zap.Logger, error) {
	cfg, isok := config.GetConfig()
	if !isok {
		return nil, err.ErrNotLoadConfig
	}

	level := getLogLevel(cfg.Log.LogLevel)

	loglevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	if config.IsConsoleLog() {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebugging := zapcore.Lock(os.Stdout)
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleDebugging, loglevel),
		)

		cl := zap.New(core)
		return cl, nil
	}

	return nil, nil
}

// InitLogger - initializes a thread-safe singleton logger
func InitLogger() (err error) {
	// once ensures the singleton is initialized only once
	onceLogger.Do(func() {
		logger, err = initLogger()
	})

	return
}

// // Log a message at the given level with given fields
// func Log(level zap.Level, message string, fields ...zap.Field) {
// 	singleton.Log(level, message, fields...)
// }

// Debug logs a debug message with the given fields
func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

// Info logs a debug message with the given fields
func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

// Warn logs a debug message with the given fields
func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

// Error logs a debug message with the given fields
func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

// Fatal logs a message than calls os.Exit(1)
func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

// SyncLogger - sync logger
func SyncLogger() {
	logger.Sync()
}

// JSON - make json to field
func JSON(key string, obj interface{}) zap.Field {
	s, err := json.Marshal(obj)
	if err != nil {
		return zap.Error(err)
	}

	return zap.String(key, string(s))
}
