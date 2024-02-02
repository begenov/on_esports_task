package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.SugaredLogger

func init() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
	zapLogger = logger.Sugar()
	defer zapLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func Debug(args ...interface{}) {
	zapLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	zapLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zapLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zapLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	zapLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	zapLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	zapLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	zapLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	zapLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	zapLogger.Fatalf(template, args...)
}
