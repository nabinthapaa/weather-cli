package structs

// Condition represents the weather condition with a descriptive text.
type Condition struct {
	Text string `json:"text"`
}

// Day represents the daily forecast data, including temperatures and UV index.
type Day struct {
	Condition Condition `json:"condition"`
	MaxTempC  float64   `json:"maxtemp_c"`
	MinTempC  float64   `json:"mintemp_c"`
	AvgTempC  float64   `json:"avgtemp_c"`
	UV        float64   `json:"uv"`
}

// ForecastDay represents the forecast data for a single day, including hourly data.
type ForecastDay struct {
	Hour []Hour `json:"hour"`       // Hourly forecast data
	Day  Day    `json:"day"`        // Daily forecast summary
	Time int64  `json:"date_epoch"` // Unix timestamp for the date
}

// Hour represents the hourly forecast data, including temperature, wind speed, and UV index.
type Hour struct {
	Condition    Condition `json:"condition"`
	Time         int64     `json:"time_epoch"`     // Unix timestamp for the hour
	TempC        float64   `json:"temp_c"`         // Temperature in Celsius
	WindSpeed    float64   `json:"wind_kph"`       // Wind speed in kilometers per hour
	ChanceOfRain int       `json:"chance_of_rain"` // Chance of rain as a percentage
	UV           float64   `json:"uv"`             // UV index
}

// Forecast represents the overall forecast data, including multiple days of forecast.
type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"` // Array of forecast days
}

// Location represents the location for which the weather data is provided.
type Location struct {
	Name    string `json:"name"`    // Name of the location
	Country string `json:"country"` // Country of the location
}

// Current represents the current weather conditions.
type Current struct {
	Condition Condition `json:"condition"` // Current weather condition
	TempC     float64   `json:"temp_c"`    // Current temperature in Celsius
	WindSpeed float64   `json:"wind_kph"`  // Current wind speed in kilometers per hour
}

// CurrentResponse represents the response containing current weather data.
type CurrentResponse struct {
	Location Location `json:"location"` // Location data
	Current  Current  `json:"current"`  // Current weather data
}

// ForecastResponse represents the response containing forecast weather data.
type ForecastResponse struct {
	Forecast Forecast `json:"forecast"` // Forecast data
	Location Location `json:"location"` // Location data
	Current  Current  `json:"current"`  // Current weather data
}
