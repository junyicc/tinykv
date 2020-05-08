package zaplog

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"
)

// logger is a global Logger
var logger *zap.Logger = newLogger()

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Flush() error {
	return logger.Sync()
}

func newLogger() *zap.Logger {
	// print all level log
	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	logPath := viper.GetString("Log.Path")
	if logPath == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		logPath = filepath.Join(home, ".tinykv", "logs", "tinykv.log")
	}
	hook := lumberjack.Logger{
		Filename: logPath,
		MaxSize:  1024,
		MaxAge:   7,
	}
	// writers
	fileWriter := zapcore.AddSync(&hook)
	consoleWriter := zapcore.Lock(os.Stdout)
	// encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, fileWriter, allPriority),
		zapcore.NewCore(consoleEncoder, consoleWriter, allPriority),
	)

	Logger := zap.New(core, zap.AddCaller())
	Logger.Info("init Logger successfully")

	return Logger
}
