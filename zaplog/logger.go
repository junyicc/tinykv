package zaplog

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	logger *zap.Logger
	w      io.Writer
}

// zaplogger is a global Logger
var zaplogger *zapLogger = newLogger()

func Debug(msg string, fields ...zap.Field) {
	zaplogger.logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zaplogger.logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zaplogger.logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zaplogger.logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	zaplogger.logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zaplogger.logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zaplogger.logger.Fatal(msg, fields...)
}

func Flush() error {
	return zaplogger.logger.Sync()
}

func Writer() io.Writer {
	return zaplogger.w
}

func newLogger() *zapLogger {
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

	w := io.MultiWriter(fileWriter, consoleWriter)
	zaplogger := &zapLogger{
		logger: zap.New(core, zap.AddCaller()),
		w:      w,
	}
	zaplogger.logger.Info("init logger successfully")

	return zaplogger
}
