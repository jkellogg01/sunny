package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiKey   string `mapstructure:"api_key"`
	HomeCity string `mapstructure:"home_city"`
}

func InitConfig() Config {
	vp := viper.New()

	vp.SetConfigName("sunny")
  vp.SetConfigType("json")
  vp.AddConfigPath(".")
	vp.AddConfigPath("$HOME/.config/sunny")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	err = vp.Unmarshal(&config)
	if err != nil {
    panic(err)
	}

	return config 
}
