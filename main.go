package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/whosoup/pegcap-bridge/api"
)

func main() {

	var pegnet, factomd string
	var port int

	flag.StringVar(&pegnet, "pegnet", "http://localhost:8070", "location of the pegnetd endpoint (no trailing slash)")
	flag.StringVar(&factomd, "factom", "http://localhost:8088/v2", "location of the factomd endpoint (with trailing /v2)")
	flag.IntVar(&port, "port", 5151, "port to serve on")
	flag.Parse()

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
	api.Init(pegnet, factomd)

	// Routes
	e.GET("/", noexist)
	e.GET("/v1/asset/names", api.AssetNames)
	e.GET("/v1/24hour/:height", api.OneDay)
	e.GET("/v1/all/:height", api.All)
	e.GET("/v1/rich/:asset", api.RichList)
	e.GET("/v1/rich", api.GlobalRichList)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Handler
func noexist(c echo.Context) error {
	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "unsupported path"}
}
