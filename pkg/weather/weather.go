package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
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

func GetWeather(lat float64, lon float64, key string) WeatherResponse {
	endpoint := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=imperial", lat, lon, key)
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("Weather API status: %v\n", res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	return weather
}
