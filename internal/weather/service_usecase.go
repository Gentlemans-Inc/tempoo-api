package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Service for weather endpoint
type Service struct{}

// GetCurrentWeather on OpenWeather API
func (s Service) GetCurrentWeather(coords *Request) (*Response, error) {
	openWeatherURI := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", coords.Latitude, coords.Longitude, os.Getenv("OPEN_WEATHER_API_KEY"))

	response, err := http.Get(openWeatherURI)

	// if request to OpenWeather is alright
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprint(response.StatusCode))
	}

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	m := &openWeatherCherryPick{}

	if err = json.Unmarshal([]byte(bodyBytes), &m); err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := new(Response)

	res.CurrentTemperature = m.Main.CurrentTemperature
	res.Description = m.Weather[0].Main
	res.Humidity = m.Main.Humidity
	res.Main = m.Weather[0].Main
	res.MaximumTemperature = m.Main.MaximumTemperature
	res.MinimumTemperature = m.Main.MinimumTemperature
	res.WindSpeed = m.Wind.Speed

	return res, nil
}
