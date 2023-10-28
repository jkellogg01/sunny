package geocoding

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Geocoding struct {
	City      string  `json:"name"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func CityGeo(city string, key string) (Geocoding, error) {
	endpoint := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&limit=1&appid=%v", city, key)
	res, err := http.Get(endpoint)
	if err != nil {
		return Geocoding{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Geocoding{}, fmt.Errorf("geocoding API status: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Geocoding{}, nil
	}

	var geocoding []Geocoding
	err = json.Unmarshal(body, &geocoding)
	if err != nil {
		return Geocoding{}, err
	}

	return geocoding[0], nil
}
