package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/whosoup/pegcap-middleware/api"
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
	e.Use(middleware.Logger())
	//e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST},
	}))

	api := new(api.Api)
	api.DB = db

	// Routes
	e.GET("/", noexist)
	e.GET("/v1/asset/list/:height", api.AssetList)
	e.GET("/v1/asset/names", api.AssetNames)
	e.GET("/v1/asset/name/:code", api.AssetName)

	e.GET("/v1/heights", api.Heights)

	// Start server
	e.Logger.Fatal(e.Start(":5151"))
}

// Handler
func noexist(c echo.Context) error {
	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "unsupported path"}
}
