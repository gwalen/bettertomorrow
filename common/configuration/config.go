package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("default")
	viper.AddConfigPath("./configfiles")
	err := viper.ReadInConfig() 
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}