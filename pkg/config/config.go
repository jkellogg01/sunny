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
	cfgPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}
	rawCfg, err := os.ReadFile(cfgPath)
	if err == os.ErrNotExist {
		mustInitConfig(cfgPath)
		rawCfg, err = os.ReadFile(cfgPath)
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
	cfgPath, err := getConfigPath()
	if err != nil {
		return err
	}
	file, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	cfgWrite, err := json.MarshalIndent(cfg, "" , "\t")
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
	cfg, err := json.MarshalIndent(cfgBlank, "" , "\t")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(cfg)
	if err != nil {
		panic(err)
	}
}

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.config/sunny/sunny.json", homePath), nil
}