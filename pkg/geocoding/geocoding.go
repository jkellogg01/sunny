package geocoding

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Geocoding struct {
	City      string  `json:"name"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func CityGeo(city string, key string) Geocoding {
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
