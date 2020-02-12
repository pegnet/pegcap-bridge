package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/labstack/echo/v4"
	"github.com/pegnet/pegnetd/srv"
)

type Api struct {
	Cli    *srv.Client
	Factom *factom.Client
}

func (api *Api) BadRequest(message string) *echo.HTTPError {
	return &echo.HTTPError{Code: http.StatusInternalServerError, Message: message}
}

func (api *Api) Init(pegnetd string, factomd string) {
	api.Cli = srv.NewClient()
	api.Cli.PegnetdServer = pegnetd
	api.Factom = &factom.Client{FactomdServer: factomd}
}

func (api *Api) GetBlockTime(height uint32) (time.Time, error) {
	dblock := new(factom.DBlock)
	dblock.Height = height
	if err := dblock.Get(context.Background(), api.Factom); err != nil {
		return time.Time{}, err
	}

	return dblock.Timestamp, nil
}

func (api *Api) GetSyncHeight() int {
	var res srv.ResultGetSyncStatus
	err := api.Cli.Request("get-sync-status", nil, &res)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(res.Sync)
}

func (api *Api) VerifyHeight(c echo.Context) (int, error) {
	h := c.Param("height")

	if h == "current" {
		cur := api.GetSyncHeight()
		if cur == -1 {
			return -1, api.BadRequest("unable to contact endpoint")
		}
		return cur, nil
	}

	hs, err := strconv.Atoi(h)
	if err != nil {
		return -1, api.BadRequest("unable to parse height")
	}

	cur := api.GetSyncHeight()
	if cur == -1 {
		return -1, api.BadRequest("unable to contact endpoint")
	}

	if hs > cur {
		return -1, api.BadRequest("data does not exist yet")
	}
	return hs, nil
}

func (api *Api) VerifyAsset(c echo.Context) (string, bool) {
	a := c.Param("asset")
	fixed := RTrans(a)
	_, ok := names[a]
	return fixed, ok
}

func Uint64ToFloat(u uint64) float64 {
	return float64(u) / 1e8
}
