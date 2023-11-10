package config

import (
	"fmt"

	"github.com/spf13/viper"
)
type HomeCity struct {
	City      string  `mapstructure:"name"`
	State     string  `mapstructure:"state"`
	Country   string  `mapstructure:"country"`
	Latitude  float64 `mapstructure:"lat"`
	Longitude float64 `mapstructure:"lon"`
}
type Config struct {
	ApiKey   string   `mapstructure:"api_key"`
	HomeCity HomeCity `mapstructure:"home_city"`
}

func ExtractConfig() (Config, error) {
	vp := viper.New()

	vp.SetConfigName("sunny")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("$HOME/.config/sunny")

	err := vp.ReadInConfig()
	switch err.(type) {
	case viper.ConfigFileAlreadyExistsError:
		fmt.Println("No config file found, initializing new config...")
		err = initConfig(vp)
		if err != nil {
			return Config{}, err
		}
	case error:
		return Config{}, err
	default:
		fmt.Println("Read config file successfully!")
	}

	var config Config
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func SetUserKey(key string) error {
	vp := viper.New()

	vp.SetConfigName("sunny")
	vp.SetConfigType("json")
	vp.AddConfigPath("$HOME/.config/sunny")

	vp.Set("api_key", key)
	return vp.WriteConfig()
}

func SetUserHome(geo HomeCity) error {
	vp := viper.New()
	fmt.Println(geo)

	vp.SetConfigName("sunny")
	vp.SetConfigType("json")
	vp.AddConfigPath("$HOME/.config/sunny")

	vp.Set("home_city", geo)
	return vp.WriteConfig()
}

// TODO: this function will create a sunny.json in the correct directory, and maybe prompt the user for an API key
func initConfig(vp *viper.Viper) error {
	err := vp.WriteConfig()
	return err
}
