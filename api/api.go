package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

type Api struct {
	DB *leveldb.DB
	C  *factom.Client
}

func (api *Api) BadRequest(message string) *echo.HTTPError {
	return &echo.HTTPError{Code: http.StatusInternalServerError, Message: message}
}

func (api *Api) Init(factomd string) {
	api.C = &factom.Client{FactomdServer: factomd}
}

func (api *Api) GetRealHeight() int {
	var h factom.Heights
	err := h.Get(api.C)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(h.DirectoryBlock)
}

func (api *Api) VerifyHeight(c echo.Context) (int, error) {
	h := c.Param("height")
	hs, err := strconv.Atoi(h)
	if err != nil {
		return -1, api.BadRequest("unable to parse height")
	}

	if hs < Genesis {
		return -1, api.BadRequest("invalid height")
	}

	cur := api.GetRealHeight()
	if cur == -1 {
		return -1, api.BadRequest("unable to contact endpoint")
	}

	if hs > cur {
		return -1, api.BadRequest("data does not exist yet")
	}
	return hs, nil
}

func Uint64ToFloat(u uint64) float64 {
	return float64(u) / 1e8
}
