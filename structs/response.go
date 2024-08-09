package structs

type Condition struct {
	Text string `json:"text"`
}

type Day struct {
	Condition Condition `json:"condition"`
	MaxTempC  float64   `json:"maxtemp_c"`
	MinTempC  float64   `json:"mintemp_c"`
	AvgTempC  float64   `json:"avgtemp_c"`
	UV        float64   `json:"uv"`
}

type ForecastDay struct {
	Hour []Hour `json:"hour"`
	Day  Day    `json:"day"`
	Time int64  `json:"date_epoch"`
}

type Hour struct {
	Condition    Condition `json:"condition"`
	Time         int64     `json:"time_epoch"`
	TempC        float64   `json:"temp_c"`
	WindSpeed    float64   `json:"wind_kph"`
	ChanceOfRain int       `json:"chance_of_rain"`
	UV           float64   `json:"uv"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	Condition Condition `json:"condition"`
	TempC     float64   `json:"temp_c"`
	WindSpeed float64   `json:"wind_kph"`
}

type CurrentResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type ForecastResponse struct {
	Forecast Forecast `json:"forecast"`
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
