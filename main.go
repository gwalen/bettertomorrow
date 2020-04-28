package main

import (
	"bettertomorrow/route"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

func MyMarshalStack(err error) interface{} {
	return fmt.Sprintf("%+v \n", err)
}

func InitZapLogger() {
	// z := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	// logger, _ := zap.NewProduction()
	// logger, _ := zap.NewDevelopment()
	// logger, _ := zap.New(zapcore.NewCore(z, )))

}

func main()  {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// zerolog.ErrorStackMarshaler = MyMarshalStack
	

	// output := zerolog.ConsoleWriter{Out: os.Stderr}
	// output.FormatFieldName = func(i interface{}) string {
		// return fmt.Sprintf("%+v \n", i)
	// }

	// log.Logger = log.Output(output)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Start application")
	router := route.Init()
	router.Logger.Fatal(router.Start(":8000"))
}
