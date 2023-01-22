package logs

import (
	"LMWN-assignment/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	if config.C.Environment == 1 {
		logConfig := zap.NewProductionConfig()
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		var err error
		log, err = logConfig.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		} else {

		}
	} else {
		logConfig := zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		var err error
		log, err = logConfig.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
