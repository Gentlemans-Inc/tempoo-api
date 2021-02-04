package weather

// Request handles data for /weather/current
type Request struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Response handles response for weather/current
type Response struct {
	Main               string  `json:"main"`
	CurrentTemperature float32 `json:"current_temperature"`
	Description        string  `json:"description"`
	Humidity           uint    `json:"humidity"`
	MinimumTemperature float32 `json:"minimum_temperature"`
	MaximumTemperature float32 `json:"maximum_temperature"`
	WindSpeed          float32 `json:"wind_speed"`
}

type openWeatherCherryPick struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		CurrentTemperature float32 `json:"temp"`
		Humidity           uint    `json:"humidity"`
		MinimumTemperature float32 `json:"temp_min"`
		MaximumTemperature float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
	}
}
