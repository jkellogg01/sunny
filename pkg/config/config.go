package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ApiKey   string `mapstructure:"api_key"`
	HomeCity string `mapstructure:"home_city"`
}

func InitConfig() Config {
	vp := viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	err = vp.Unmarshal(&config)
	if err != nil {
    log.Panic(err)
	}

	return config 
}
