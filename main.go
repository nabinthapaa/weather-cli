package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"web-scrapper/config"
	"web-scrapper/structs"

	"github.com/alecthomas/kong"
	"github.com/jedib0t/go-pretty/v6/table"
)

type CLI struct {
	Location string `help:"Location whose weather is to be fetched" default:"Kathmandu"`
	Type     string `help:"Current or Forecast" default:"forecast"`
	Format   string `help:"Daily or Hourly" default:"hourly"`
	NoOfDays string `help:"No. of days for which forecast is needed" default:"10"`
}

func main() {
	var CLI CLI
	_ = kong.Parse(&CLI)
	env, err := config.GetConfig("./.env")
	if err != nil {
		panic(err)
	}

	res, err := http.Get("http://api.weatherapi.com/v1/" + CLI.Type + ".json?q=" + CLI.Location + "&key=" + env.WeatherApi + "&days=" + CLI.NoOfDays)
	if err != nil {
		fmt.Println("Weather api unavailble")
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Weather api unavailble")
		panic(err)
	}

	t := table.NewWriter()
	t.SetCaption("Weather Report")
	switch CLI.Format {
	case "daily":
		switch CLI.Type {
		case "forecast":
			printForecast(t, body)
		case "current":
			printCurrent(t, body)
		default:
			fmt.Println("Invalid arguement")
		}
	case "hourly":
		printHourly(t, body)
	}

	fmt.Println(t.Render())
}

func printForecast(t table.Writer, body []byte) {
	var data structs.ForecastResponse
	json.Unmarshal(body, &data)
	t.AppendHeader(table.Row{"Date", "Max Temp", "Avg Temp", "Min Temp", "Condition"})
	for _, forecast := range data.Forecast.ForecastDay {
		date := time.Unix(forecast.Time, 0).Format("2006-01-02")
		t.AppendRow(
			table.Row{
				date,
				forecast.Day.MaxTempC,
				forecast.Day.AvgTempC,
				forecast.Day.MinTempC,
				forecast.Day.Condition.Text,
			},
		)
	}
}

func printCurrent(t table.Writer, body []byte) {
	var data structs.CurrentResponse
	json.Unmarshal(body, &data)
	t.AppendHeader(table.Row{"Location", "Temperature", "Wind Speed", "Condition"})
	t.AppendRow(table.Row{
		data.Location.Name + " - " + data.Location.Country,
		data.Current.TempC,
		data.Current.WindSpeed,
		data.Current.Condition.Text,
	})
}

func printHourly(t table.Writer, body []byte) {
	var data structs.ForecastResponse
	json.Unmarshal(body, &data)
	t.AppendHeader(table.Row{"Time", "Chance of Rain", "Temperature", "Wind Speed", "UV", "Condition"})
	for _, forecast := range data.Forecast.ForecastDay[0].Hour {
		date := time.Unix(forecast.Time, 0)
		if date.Before(time.Now()) {
			continue
		}
		t.AppendRow(
			table.Row{
				date.Format("03:04 PM"),
				fmt.Sprintf("\t%d%%", forecast.ChanceOfRain),
				fmt.Sprintf("\t%.2f°C", forecast.TempC),
				fmt.Sprintf("\t%.2f KPH", forecast.WindSpeed),
				forecast.UV,
				forecast.Condition.Text,
			},
		)
	}
}
