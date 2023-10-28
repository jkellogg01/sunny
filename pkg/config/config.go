package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiKey   string `mapstructure:"api_key"`
	HomeCity string `mapstructure:"home_city"`
}

func InitConfig() (Config, error) {
	vp := viper.New()

	vp.SetConfigName("sunny")
  vp.SetConfigType("json")
  vp.AddConfigPath(".")
	vp.AddConfigPath("$HOME/.config/sunny")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
