package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/model"
	"io/ioutil"
	"net/http"
	"os"
)

// Service for weather endpoint
type Service struct{}

// GetCurrentWeather on OpenWeather API
func (s Service) GetCurrentWeather(coords *model.Request) (*model.Response, error) {
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

	m := &model.Forecast{}

	if err = json.Unmarshal(bodyBytes, &m); err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := &model.Response{}

	res.ParseFromForecast(m)

	return res, nil
}
