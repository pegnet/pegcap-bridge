package api

import (
	"encoding/json"
	"net/http"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/labstack/echo/v4"
)

var Genesis = 206422
var V2 = 210330

func (api *Api) Heights(c echo.Context) error {

	h := api.GetRealHeight()
	if h == -1 {
		return api.BadRequest("unable to contact endpoint")
	}

	heights := make(map[string]int)
	heights["Current"] = h
	heights["Genesis"] = Genesis

	var dblock factom.DBlock
	dblock.Header.Height = uint32(h)
	dblock.Get(api.C)

	if !dblock.IsPopulated() {
		return api.BadRequest("unable to contact endpoint")
	}

	heights["Blocktime"] = int(dblock.Header.Timestamp.Unix())

	js, _ := json.Marshal(heights)
	return c.JSONBlob(http.StatusOK, js)
}

func (api *Api) Height(c echo.Context) error {
	h, err := api.VerifyHeight(c)
	if err != nil {
		return err
	}
	var dblock factom.DBlock
	dblock.Header.Height = uint32(h)
	dblock.Get(api.C)

	if !dblock.IsPopulated() {
		return api.BadRequest("unable to contact endpoint")
	}

	height := make(map[string]int)
	height["Blocktime"] = int(dblock.Header.Timestamp.Unix())

	js, _ := json.Marshal(height)
	return c.JSONBlob(http.StatusOK, js)
}
