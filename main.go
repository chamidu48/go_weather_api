package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var client *http.Client

// --to get country's lat an lon--
type Country struct {
	Name string  `json:"name"`
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
}

// --to get temp details--
type Weather struct {
	Main TempDetails `json:"main"`
}
type TempDetails struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

// --get JSON data
func getJSON(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

// --get weather details---
func getWeather(c echo.Context) error {
	country := c.QueryParam("country")

	apiKey := "376960a6acc4f79ad0baafc6d71b266d"
	url := "http://api.openweathermap.org/geo/1.0/direct?q=" + country + "&limit=5&appid=" + apiKey

	var countries []Country //--store country details--
	var weather Weather     //--store weather details of a country--

	//--get lat and lon--
	err := getJSON(url, &countries)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("Error: %s", err.Error()))
	}

	if len(countries) > 0 {
		//--recieve lat/lon details--
		lat := fmt.Sprintf("%f", countries[0].Lat)
		lon := fmt.Sprintf("%f", countries[0].Lon)

		//--get weather details from api--
		err1 := getJSON("https://api.openweathermap.org/data/2.5/weather?lat="+lat+"&lon="+lon+"&appid="+apiKey, &weather)

		if err1 != nil {
			return c.String(http.StatusOK, fmt.Sprintf("Error: %s", err1.Error()))
		} else {
			return c.String(http.StatusOK, fmt.Sprintf("%s's Weather Details\nTemperature : %f\nPressure : %d\nHumidity : %d",
				country, weather.Main.Temp, weather.Main.Pressure, weather.Main.Humidity))
		}
	}

	return c.String(http.StatusOK, "No location found")
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/weather", getWeather)

	e.Logger.Fatal(e.Start(":1323"))
}
