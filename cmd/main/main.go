package main

import (
	"flag"
	"fmt"
	"log"
	"sunny/pkg/geocoding"
	"sunny/pkg/weather"
  "sunny/pkg/config"
)

func main() {
	config, err := config.InitConfig();
	if err != nil {
		log.Fatal(err)
	}
	home, key := config.HomeCity, config.ApiKey

	userCity := flag.String("c", home, "Enter the city where you would like to look up the weather")
	flag.Parse()
	geo := geocoding.CityGeo(*userCity, key)
	city, state, country, lat, lon := geo.City, geo.State, geo.Country, geo.Latitude, geo.Longitude
  currentWeather := weather.GetWeather(lat, lon, key);
	temp, feels, desc, humidity := currentWeather.Main.Temperature, currentWeather.Main.FeelsLike, currentWeather.Weather[0].Description, currentWeather.Main.Humidity

	fmt.Printf("%s, %s, %s\nlat: %f lon: %f\n", city, state, country, lat, lon)
	fmt.Printf("Currently %.2fºF and %s\nFeels like %.2fºF, %v%% Humidity\n", temp, desc, feels, humidity)
}
