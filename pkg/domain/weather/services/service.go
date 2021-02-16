package services

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/model"
)

// WeatherService for weather endpoint
type WeatherService interface {
	GetCurrentWeather(coords *model.Request) (*model.Response, error)
}

// NewWeatherService for weather handler
func NewWeatherService() (service WeatherService) {
	service = &Service{}
	return
}
