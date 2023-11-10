package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jkellogg01/sunny/pkg/geocoding"
)

const cfg_path string = "$HOME/.config/sunny.json"
type Config struct {
	ApiKey   string              `json:"api_key"`
	HomeCity geocoding.Geocoding `json:"home_city"`
}

func ExtractConfig() (Config, error) {
	rawCfg, err := os.ReadFile(cfg_path)
	if err == os.ErrNotExist {
		mustInitConfig(cfg_path)
		rawCfg, err = os.ReadFile(cfg_path)
	}
	if err != nil {
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

func (cfg *Config) UpdateConfig() error {
	file, err := os.Open(cfg_path)
	if err != nil {
		return err
	}
	cfgWrite, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = file.Write(cfgWrite)
	return err
}

func mustInitConfig(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	_, err = file.Write(cfg)
	if err != nil {
		panic(err)
	}
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
