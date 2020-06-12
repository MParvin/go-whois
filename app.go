package main

// go get github.com/labstack/echo/v4
// go get github.com/likexian/whois-go

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/likexian/whois-go"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/:queryString", getWhois)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func getWhois(c echo.Context) error {
	var domainOrIP = c.Param("queryString")
	whoisData, err := whois.Whois(domainOrIP)
	if err == nil {
		return c.String(http.StatusOK, whoisData)
	}
	return c.String(http.StatusBadRequest, err.Error())
}
