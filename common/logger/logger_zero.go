package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

/**
 * basic implmentation for real production use it had to enhanced
 */

type LoggerZero struct {
	loggerType string //TODO: enum
	caller string
	logger zerolog.Logger
}

func (lz *LoggerZero) ForDev() {
	lz.logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	lz.loggerType = "dev"
	lz.caller = callerForLogger()
}

func (lz *LoggerZero) ForProd() {
	// lz.logger = log.Logger.With().Caller().Logger() // cant use caller as we have extra facade layer for logger (so we need to get one frame deeper than defualt impl)
	lz.logger = log.Logger
	lz.loggerType = "prod"
	lz.caller = callerForLogger()
}

func (lz *LoggerZero) InitZeroLogger() *LoggerZero {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	lz.ForProd() 	// defalut type = prod with json

	return lz
}

func (lz *LoggerZero) Error(msg string, err error) {
	if lz.loggerType == "dev" { 
		lz.logger.Err(err).Str("caller", lz.caller).Msgf("%v %+v", msg, err)  
	} else {
		lz.logger.Error().Stack().Err(err).Str("caller", lz.caller).Msg(msg)
	}
}

func (lz *LoggerZero) Warn(msg string) {
	lz.logger.Warn().Str("caller", lz.caller).Msg(msg)		
}

func (lz *LoggerZero) Info(msg string) {
	lz.logger.Info().Str("caller", lz.caller).Msg(msg)
}

func (lz *LoggerZero) Debug(msg string) {
	lz.logger.Debug().Str("caller", lz.caller).Msg(msg)	
}