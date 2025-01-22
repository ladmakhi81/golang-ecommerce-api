package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() ILogger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(getColoredLevel(level))
		},
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	return ZapLogger{
		logger: zap.New(core),
	}
}

func (zapLogger ZapLogger) Info(msg string) {
	zapLogger.logger.Info(msg)
}

func (zapLogger ZapLogger) InfoWithMeta(msg string, meta any) {
	zapLogger.logger.Info(msg, zap.Any("meta", meta))
}

func getColoredLevel(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return "\033[36mDEBUG\033[0m" // Cyan
	case zapcore.InfoLevel:
		return "\033[32mINFO\033[0m" // Green
	case zapcore.WarnLevel:
		return "\033[33mWARN\033[0m" // Yellow
	case zapcore.ErrorLevel:
		return "\033[31mERROR\033[0m" // Red
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return "\033[35mFATAL\033[0m" // Magenta
	default:
		return level.String()
	}
}
