package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getWeather(c echo.Context) error {
	country := c.QueryParam("country")

	apiKey := "376960a6acc4f79ad0baafc6d71b266d"
	apiGetCord := "http://api.openweathermap.org/geo/1.0/direct?q=" + country + "&limit=5&appid=" + apiKey

	// lat := c.QueryParam("lat")
	// lon := c.QueryParam("lon")

	return c.String(http.StatusOK, fmt.Sprintf("Your location's weather is : %s", apiGetCord))
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/weather", getWeather)

	e.Logger.Fatal(e.Start(":1323"))
}
