package config

import (
  "github.com/spf13/viper"
)

type Config struct {
	ApiKey 		string `mapstructure:"api_key"`
	HomeCity 	string `mapstructure:"home_city"`
}

func InitConfig() (Config, error) {
	vp := viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
