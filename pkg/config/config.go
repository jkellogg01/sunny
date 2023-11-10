package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ApiKey   string `json:"api_key"`
	HomeCity struct {
		City      string  `json:"name"`
		State     string  `json:"state"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	} `json:"home_city"`
}

func ExtractConfig() (Config, error) {
	rawCfg, err := os.ReadFile("$HOME/.config/sunny")
	if err == os.ErrNotExist {
		// create a new config file
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
