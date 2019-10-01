package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/whosoup/pegcap-bridge/api"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	userPath := u.HomeDir
	dbPath := userPath + "/.pegcap"
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Could not create the directory %s", dbPath))
	}

	db, err := leveldb.OpenFile(dbPath+"/dummy", nil)
	if err != nil {
		panic(fmt.Sprintln("Could not create db:", err))
	}

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
	api.Init("https://api.factomd.net")
	api.DB = db

	// Routes
	e.GET("/", noexist)
	e.GET("/v1/asset/list/:height", api.AssetList)
	e.GET("/v1/asset/names", api.AssetNames)
	e.GET("/v1/asset/name/:code", api.AssetName)

	e.GET("/v1/heights", api.Heights)
	e.GET("/v1/height/:height", api.Height)
	e.GET("/v1/24hour/:height", api.OneDay)
	e.GET("/v1/rates/:height", api.Rates)
	e.GET("/v1/market/:height", api.Market)
	e.GET("/v1/all/:height", api.All)

	// Start server
	e.Logger.Fatal(e.Start(":5151"))
}

// Handler
func noexist(c echo.Context) error {
	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "unsupported path"}
}
