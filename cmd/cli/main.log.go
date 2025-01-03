package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
) // logs by zap (uber)

func main() {
	// var sugar = zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "DaoHai", 25) // like fmt.Printf(format, a)
	// // result:  {"level":"info","msg":"Hello name:DaoHai, age:25"}

	// //logger
	// var logger = zap.NewExample()
	// logger.Info("Hello", zap.String("name", "DaoHai"), zap.Int("age", 25))
	// // result:  {"level":"info","msg":"Hello","name":"DaoHai","age":25}


	// // example
	// var loggerExample = zap.NewExample()
	// loggerExample.Info("Hello NewExample")
	// // result:   {"level":"info","msg":"Hello NewExample"}

	// // Development
	// var loggerDevelopment, _ = zap.NewDevelopment()
	// loggerDevelopment.Info("Hello NewDevelopment")
	// // result:   2025-01-03T22:19:24.898+0700    INFO    cli/main.log.go:24      Hello NewDevelopment

	// //Production
	// var loggerProduction, _ = zap.NewProduction()
	// loggerProduction.Info("Hello NewProduction")
	// // result:   {"level":"info","ts":1735917564.8992972,"caller":"cli/main.log.go:28","msg":"Hello NewProduction"}

	var encoder = getEncoderLog()
	var sync = getWriterSync()
	var core = zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	var logger = zap.New(core)

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	var file, _ = os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	var syncFile = zapcore.AddSync(file)
	var syncConsole = zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}