package api

import (
	"encoding/json"
	"net/http"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/labstack/echo/v4"
)

var Genesis = 206422

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
