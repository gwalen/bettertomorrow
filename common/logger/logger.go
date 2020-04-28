package logger

import (
	"path/filepath"
	"runtime"
)

var loggerFacade LoggerFacade

//TODO: use config from file
const (
	zapLogger  = "zap"
	zeroLogger = "zero"

	loggerTypeDev = "dev"
	loggerTypeProd = "prod"
)


type Logger interface {
	Error(msg string, err error)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
}

type LoggerFacade struct {
	logger Logger
}


// question is weather zap/zero performs better with just one instance (singleton) or with many
// TODO: use config file
func ProvideLogger(loggerVendor string, loggerType string) Logger {
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

func callerForLogger() string {
	_, filePath, _, ok := runtime.Caller(3)
	dir, file := filepath.Split(filePath)
    caller := filepath.Join(filepath.Base(dir), file)
    if !ok {
        return "<unknown>"
	}
	return caller
}