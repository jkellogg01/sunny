package geocoding

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Geocoding struct {
	City      string  `json:"name"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

const max_cities int = 10

func CityGeo(city string, key string) ([]Geocoding, error) {
	endpoint := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%v&appid=%v&limit=%v", city, key, max_cities)
	res, err := http.Get(endpoint)
	if err != nil {
		return []Geocoding{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []Geocoding{}, fmt.Errorf("geocoding API status: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Geocoding{}, nil
	}

	var geocoding []Geocoding
	err = json.Unmarshal(body, &geocoding)
	if err != nil {
		return []Geocoding{}, err
	}

	return geocoding, nil
}

func HandleGeoCollision(geos []Geocoding) (Geocoding, error) {
    fmt.Println("Found multiple cities with this name.")
    for i, geo := range geos {
        fmt.Printf("%d. %v, %v, %v", i + 1, geo.City, geo.State, geo.Country)
    }
    fmt.Println("Which would you like to choose?")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadBytes('\n')
    if err != nil {
        return Geocoding{}, err
    }
    menuItem, err := strconv.Atoi(string(input))
    if err != nil {
        return Geocoding{}, err
    }
    if menuItem > len(geos) || menuItem < 1 {
        return Geocoding{}, fmt.Errorf("selection outside of range")
    }
    return geos[menuItem - 1], nil
}