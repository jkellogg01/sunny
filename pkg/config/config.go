package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jkellogg01/sunny/pkg/geocoding"
)

type Config struct {
	ApiKey   string              `json:"api_key"`
	HomeCity geocoding.Geocoding `json:"home_city"`
}

func ExtractConfig() (Config, error) {
	rawCfg, err := os.ReadFile("$HOME/.config/sunny")
	if err == os.ErrNotExist {
		initConfig("$HOME/.config/sunny")
	} else if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(rawCfg, &cfg)
	if err != nil {
		return Config{}, err
	}

	fmt.Println(cfg)
	return cfg, nil
}

func initConfig(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	cfgBlank := Config{
		ApiKey: "",
		HomeCity: geocoding.Geocoding{
			City:      "",
			State:     "",
			Country:   "",
			Latitude:  0,
			Longitude: 0,
		},
	}
	cfg, err := json.Marshal(cfgBlank)
	if err != nil {
		return err
	}
	file.Write(cfg)
	return nil
}

// func ExtractConfig() (Config, error) {
// 	vp := viper.New()

// 	vp.SetConfigName("sunny")
// 	vp.SetConfigType("json")
// 	vp.AddConfigPath(".")
// 	vp.AddConfigPath("$HOME/.config/sunny")

// 	err := vp.ReadInConfig()
// 	switch err.(type) {
// 	case viper.ConfigFileAlreadyExistsError:
// 		fmt.Println("No config file found, initializing new config...")
// 		err = initConfig(vp)
// 		if err != nil {
// 			return Config{}, err
// 		}
// 	case error:
// 		return Config{}, err
// 	default:
// 		fmt.Println("Read config file successfully!")
// 	}

// 	var config Config
// 	err = vp.Unmarshal(&config)
// 	if err != nil {
// 		return Config{}, err
// 	}

// 	return config, nil
// }

// func SetUserKey(key string) error {
// 	vp := viper.New()

// 	vp.SetConfigName("sunny")
// 	vp.SetConfigType("json")
// 	vp.AddConfigPath("$HOME/.config/sunny")

// 	vp.Set("api_key", key)
// 	return vp.WriteConfig()
// }

// func SetUserHome(geo HomeCity) error {
// 	vp := viper.New()
// 	fmt.Println(geo)

// 	vp.SetConfigName("sunny")
// 	vp.SetConfigType("json")
// 	vp.AddConfigPath("$HOME/.config/sunny")

// 	vp.Set("home_city", geo)
// 	return vp.WriteConfig()
// }

// // TODO: this function will create a sunny.json in the correct directory, and maybe prompt the user for an API key
// func initConfig(vp *viper.Viper) error {
// 	err := vp.WriteConfig()
// 	return err
// }
