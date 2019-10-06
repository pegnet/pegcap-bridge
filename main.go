package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/whosoup/pegcap-bridge/api"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} method=${method}, status=${status}, uri=${uri}, error=${error}\n",
	}))
	//e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST},
	}))

	api := new(api.Api)
	api.Init("http://localhost:8070", "http://spoon:8088/v2")

	// Routes
	e.GET("/", noexist)
	e.GET("/v1/asset/names", api.AssetNames)
	e.GET("/v1/24hour/:height", api.OneDay)
	e.GET("/v1/all/:height", api.All)

	// Start server
	e.Logger.Fatal(e.Start(":5151"))
}

// Handler
func noexist(c echo.Context) error {
	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "unsupported path"}
}
