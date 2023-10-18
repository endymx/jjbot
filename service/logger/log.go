package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "JJBot.log",
		MaxSize:    3,
		MaxBackups: 5,
		MaxAge:     7,
		LocalTime:  true,
		Compress:   false,
	}

	syncFile := zapcore.AddSync(lumberJackLogger)
	syncConsole := zapcore.AddSync(os.Stderr)
	writeSyncer := zap.CombineWriteSyncers(syncFile, syncConsole)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
	defer func(SugarLogger *zap.SugaredLogger) {
		_ = SugarLogger.Sync()
	}(SugarLogger)
}
