package main

import (
	"flag"
	"fmt"

	"github.com/jkellogg01/sunny/pkg/config"
	"github.com/jkellogg01/sunny/pkg/geocoding"
	"github.com/jkellogg01/sunny/pkg/weather"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	home, key := config.HomeCity, config.ApiKey

	userCity := flag.String("c", home, "Enter the city where you would like to look up the weather")
	flag.Parse()

	geo, err := geocoding.CityGeo(*userCity, key)
	if err != nil {
		panic(err)
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
