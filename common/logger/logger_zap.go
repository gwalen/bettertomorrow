package logger

import (
	"go.uber.org/zap"
)

/**
 * basic implmentation for real production use in had to enhanced
 */

type LoggerZap struct {
	loggerType string //TODO: enum
	logger     *zap.Logger
	caller     string
}

func (lz *LoggerZap) InitZapLogger() *LoggerZap {
	lz.ForProd() // defalut type = prod with json
	return lz
}

func (lz *LoggerZap) ForDev() {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.DisableCaller = true
	loggerConfig.DisableStacktrace = true
	logger, _ := loggerConfig.Build()
	lz.logger = logger
	lz.caller = callerForLogger()
}

func (lz *LoggerZap) ForProd() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.DisableCaller = true     // cant use caller as we have extra facade layer for logger (so we need to get one frame deeper than defualt impl)
	loggerConfig.DisableStacktrace = true // we disbale it becuase zap logs a place from which logging was called - and we want a stack trace taken from error
	logger, _ := loggerConfig.Build()
	lz.logger = logger
	lz.caller = callerForLogger()
}

func (lz *LoggerZap) Error(msg string, err error) {
	lz.logger.Sugar().With("caller", lz.caller).Errorf("%+v \n", err)
}

func (lz *LoggerZap) Warn(msg string) {
	lz.logger.Sugar().With("caller", lz.caller).Warn(msg)
}

func (lz *LoggerZap) Info(msg string) {
	lz.logger.Sugar().With("caller", lz.caller).Info(msg)
}

func (lz *LoggerZap) Debug(msg string) {
	lz.logger.Sugar().With("caller", lz.caller).Debug()
}
