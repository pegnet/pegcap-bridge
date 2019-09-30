package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (api *Api) OneDay(c echo.Context) error {
	h, err := api.VerifyHeight(c)
	if err != nil {
		return err
	}

	h -= 144
	if h < Genesis {
		return api.BadRequest("missing data")
	}
	js, _ := json.Marshal(h)
	return c.JSONBlob(http.StatusOK, js)
}
