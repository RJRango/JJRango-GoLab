package models

type CurrentWeatherResponse struct {
	City        string   `json:"city"`
	Condition   []string `json:"conditions"`
	Temperature float64  `json:"temperature"`
	Overview    string   `json:"overview"`
}

type OpenWeatherCurrentWeatherResponse struct {
	Weather `json:"weather"`
	Main    `json:"main"`
	Wind    `json:"wind"`
	Name    string `json:"name"`
}

type Main struct {
	Temp       float64 `json:"temp"`
	Feels_Like float64 `json:"feels_like"`
	Temp_Min   float64 `json:"temp_min"`
	Temp_Max   float64 `json:"temp_max"`
}

type Weather []struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}
