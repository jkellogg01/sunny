package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Geocoding struct {
	City 			string 	`json:"name"`
	State 		string 	`json:"state"`
	Country 	string 	`json:"country"`
	Latitude 	float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type WeatherResponse struct {
	// Location 			string 	`json:"name"`
	Weather 			[]struct {
		Short				string 	`json:"main"`
		Description string 	`json:"description"`
	} 										`json:"weather"`
	Main 					struct {
		Temperature float32 `json:"temp"`
		FeelsLike 	float32 `json:"feels_like"`
		Humidity 		int 		`json:"humidity"`
	} 										`json:"main"`
	Wind 					struct {
		Speed 			float32 `json:"speed"`
		Direction 	int 		`json:"deg"`
	} 										`json:"wind"`
}

type Config struct {
	ApiKey 		string `mapstructure:"api_key"`
	HomeCity 	string `mapstructure:"home_city"`
}

func cityGeo(city string, key string) Geocoding {
	endpoint := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, key)
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Geocoding API status: %d\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var geocoding []Geocoding
	err = json.Unmarshal(body, &geocoding)
	if err != nil {
		log.Fatal(err)
	}
	return geocoding[0]
}

func getWeather(lat float64, lon float64, key string) WeatherResponse {
	endpoint := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=imperial", lat, lon, key)
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Weather API status: %d\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	return weather
}

func config() (Config, error) {
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

func main() {
	config, err := config()
	if err != nil {
		log.Fatal(err)
	}
	home, key := config.HomeCity, config.ApiKey

	userCity := flag.String("c", home, "Enter the city where you would like to look up the weather")
	flag.Parse()
	geo := cityGeo(*userCity, key)
	city, state, country, lat, lon := geo.City, geo.State, geo.Country, geo.Latitude, geo.Longitude
	weather := getWeather(lat, lon, key)
	temp, feels, desc, humidity := weather.Main.Temperature, weather.Main.FeelsLike, weather.Weather[0].Description, weather.Main.Humidity

	// fmt.Println(geo)
	// fmt.Println(weather)
	fmt.Printf("%s, %s, %s\nlat: %f lon: %f\n", city, state, country, lat, lon)
	fmt.Printf("Currently %.2fºF and %s\nFeels like %.2fºF, %v%% Humidity\n", temp, desc, feels, humidity)
}