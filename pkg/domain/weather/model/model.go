package model

// Request handles data for /weather/current
type Request struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Response handles response for weather/current
type Response struct {
	Main               string  `json:"main"`
	DescriptionCode    uint    `json:"description_code"`
	CurrentTemperature float32 `json:"current_temperature"`
	Description        string  `json:"description"`
	Humidity           uint    `json:"humidity"`
	MinimumTemperature float32 `json:"minimum_temperature"`
	MaximumTemperature float32 `json:"maximum_temperature"`
	WindSpeed          float32 `json:"wind_speed"`
}

// Forecast struct model
type Forecast struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		ID          uint   `json:"id"`
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

// ParseFromForecast model to Response
func (r *Response) ParseFromForecast(f *Forecast) {
	r.CurrentTemperature = f.Main.CurrentTemperature
	r.DescriptionCode = f.Weather[0].ID
	r.Description = f.Weather[0].Main
	r.Humidity = f.Main.Humidity
	r.Main = f.Weather[0].Main
	r.MaximumTemperature = f.Main.MaximumTemperature
	r.MinimumTemperature = f.Main.MinimumTemperature
	r.WindSpeed = f.Wind.Speed
}
