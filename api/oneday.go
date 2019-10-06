package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (api *Api) OneDay(c echo.Context) error {
	h, err := api.VerifyHeight(c)
	if err != nil {
		return err
	}

	target, err := api.GetBlockTime(uint32(h))
	if err != nil {
		return api.BadRequest("unable to contact endpoint: " + err.Error())
	}
	h -= 144

	limit := 32
	for limit > 0 {
		t, err := api.GetBlockTime(uint32(h))
		if err != nil {
			return api.BadRequest("unable to contact endpoint: " + err.Error())
		}
		if !target.Before(t.Add(time.Hour * 24)) {
			js, _ := json.Marshal(h)
			return c.JSONBlob(http.StatusOK, js)
		}
		limit--
		h--
	}

	return api.BadRequest("unable to contact endpoint: " + err.Error())
}
