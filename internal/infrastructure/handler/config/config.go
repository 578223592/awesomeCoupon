package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Viper = viper.New()
)

// Init unit test ,confPath set ""
func Init(confFilePath string) {
	if confFilePath == "" {
		confFilePath = "./config/app.dev.toml"
	}

	Viper.SetConfigFile(confFilePath)

	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
