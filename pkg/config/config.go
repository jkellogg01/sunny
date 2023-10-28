package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ApiKey   string `mapstructure:"api_key"`
	HomeCity string `mapstructure:"home_city"`
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

// TODO: this function will create a sunny.json in the correct directory, and maybe prompt the user for an API key
func initConfig(vp *viper.Viper) error {
	return fmt.Errorf("functionality to initialize config file is not ready yet")
}