package logger

import (
	"go-ecommerce-backend-api/packages/setting"
	"os"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	var logLevel = config.Log_level // debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	var encoder = getEncoderLog()
	var hook = lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age, //days
		Compress:   config.Compress, // disabled by default
	}
	var core = zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), level)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}

}

func getEncoderLog() zapcore.Encoder {
	var encodeConfig = zap.NewProductionEncoderConfig()

	// 1735917564.8992972 -> 2025-01-03T22:19:24.898+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller":"cli/main.log.go:28"
	// encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}