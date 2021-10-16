package main

import (
	"net/http"

	"github.com/insights/task/database"
	"github.com/insights/task/geocoding"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/persistData", func(c echo.Context) error {
		err := database.PersistCovidData()
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, "Data persisted successfully in DB")
	})
	e.GET("/caseCount", func(c echo.Context) error {
		lat := c.QueryParam("lat")
		lon := c.QueryParam("lon")
		if lat == "" || lon == "" {
			return c.String(http.StatusBadRequest, "Provide values for lat and lon")
		}
		res, err := geocoding.GetStateFromGPSCoordinates(lat, lon)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":1122"))
	// https://app.swaggerhub.com/apis/Saksham1/Covid_API/1.0.0-oas3#/
}
