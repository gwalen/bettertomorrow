package logger

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const (
	zapLogger  = "zap"
	zeroLogger = "zero"

	loggerTypeDev = "dev"
	loggerTypeProd = "prod"
)

var loggerVendor = viper.GetString("logger.vendor")
var loggerType = viper.GetString("logger.type")


type Logger interface {
	Error(msg string, err error)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
}

// question is weather zap/zero performs better with just one instance (singleton) or with many
func ProvideLogger() Logger {
	switch loggerVendor {
	case zapLogger:
		zapInstance := LoggerZap{}
		zapInstance.InitZapLogger()
		if(loggerType == loggerTypeDev) { zapInstance.ForDev() } else { zapInstance.ForProd() }
		return &zapInstance
	case zeroLogger:
		zeroInstance := LoggerZero{}
		zeroInstance.InitZeroLogger()
		if(loggerType == loggerTypeDev) { zeroInstance.ForDev() } else { zeroInstance.ForProd() }
		return &zeroInstance
	}
	return nil
}

func ProvideDefaultLogger() Logger {
	logger := LoggerZero{}
	logger.InitZeroLogger()
	logger.ForDev()
	return &logger
}

func callerForLogger() string {
	_, filePath, _, ok := runtime.Caller(3)
	dir, file := filepath.Split(filePath)
    caller := filepath.Join(filepath.Base(dir), file)
    if !ok {
        return "<unknown>"
	}
	return caller
}