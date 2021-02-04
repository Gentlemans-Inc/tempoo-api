package weather

// WeatherService for weather endpoint
type WeatherService interface {
	GetCurrentWeather(coords *Request) (*Response, error)
}

// NewWeatherService for weather handler
func NewWeatherService() (service WeatherService) {
	service = &Service{}
	return
}
