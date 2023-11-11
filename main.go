package main

import (
	"flag"
	"fmt"

	"github.com/jkellogg01/sunny/pkg/config"
	"github.com/jkellogg01/sunny/pkg/geocoding"
	"github.com/jkellogg01/sunny/pkg/weather"
)

func main() {
	userConfig, err := config.ExtractConfig()
	if err != nil {
		panic(err)
	}
	key := userConfig.ApiKey
	home := userConfig.HomeCity

	userCity := flag.String(
		"c",
		fmt.Sprintf("%s,%s,%s", home.City, home.State, home.Country),
		"Enter the city where you would like to look up the weather",
	)
	flagKey := flag.String("k", "", "Enter your API key for the OpenWeatherMap API")
	setHome := flag.Bool("h", false, "set the city specified with the -c flag as your home")
	flag.Parse()

	// Doing this instead of using a default value so that running 'sunny --help' doesn't expose the user's api key
	if key == "" {
		if *flagKey == "" {
			fmt.Printf("It looks like you don't have an API key saved.\nIf you have an API key with openweathermap.org, run 'sunny -k {YOUR_API_KEY}' to save your API key.\nIf you need an API key, visit https://home.openweathermap.org/users/sign_up and create an account.\nOpen weather map provides a free API key that Sunny uses whenever it makes API calls.\n")
			panic(fmt.Errorf("must use an API key"))
		}
		userConfig = config.Config{
			ApiKey:   *flagKey,
			HomeCity: userConfig.HomeCity,
		}
		userConfig.UpdateConfig()
		if err != nil {
			panic(err)
		}
	}

	var geo geocoding.Geocoding
	if *userCity != "" {
		geocodings, err := geocoding.CityGeo(*userCity, key)
		if err != nil {
			panic(err)
		}

		if len(geocodings) > 1 {
			geo, err = geocoding.HandleGeoCollision(geocodings)
			if err != nil {
				panic(err)
			}
		} else {
			geo = geocodings[0]
		}
	} else {
		geo = home
	}

	if *setHome {
		userConfig = config.Config{
			HomeCity: geo,
			ApiKey:   userConfig.ApiKey,
		}
		err := userConfig.UpdateConfig()
		if err != nil {
			panic(err)
		}
	}

	city, state, country, lat, lon := geo.City, geo.State, geo.Country, geo.Latitude, geo.Longitude
	currentWeather, err := weather.GetWeather(lat, lon, key)
	if err != nil {
		panic(err)
	}
	temp, feels, desc, humidity := currentWeather.Main.Temperature, currentWeather.Main.FeelsLike, currentWeather.Weather[0].Description, currentWeather.Main.Humidity

	fmt.Printf("%v, %v, %v\nlat: %v lon: %v\n", city, state, country, lat, lon)
	fmt.Printf("Currently %.2fºF and %v\nFeels like %.2fºF, %v%% Humidity\n", temp, desc, feels, humidity)
}
