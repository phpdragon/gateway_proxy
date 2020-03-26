package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// error logger
var errorLogger *zap.Logger

func init() {
	//TODO:从配置中获取
	fileName := "zap.log"

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1024, // megabytes
		MaxBackups: 3,
		MaxAge:     7, //days
		LocalTime:  true,
		Compress:   true,
	})

	//encoder := zap.NewProductionEncoderConfig()
	//encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	productEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		// 打印在控制台
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(zap.DebugLevel)),
		// 打印在文件中,仅打印Info级别以上的日志
		zapcore.NewCore(productEncoder, fileWriter, zap.NewAtomicLevelAt(zap.InfoLevel)),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger
	defer errorLogger.Sync()
}

func GetAtomicLevel() zap.AtomicLevel {
	return zap.NewAtomicLevel()
}

func Debug(msg string, fields ...zap.Field) {
	errorLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	errorLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	errorLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	errorLogger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	errorLogger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	errorLogger.DPanic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	errorLogger.Fatal(msg, fields...)
}