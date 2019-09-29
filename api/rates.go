package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) Rates(c echo.Context) error {
	h, err := a.VerifyHeight(c)
	if err != nil {
		return err
	}
	a.Generate(h)
	if Outage[h] {
		return a.BadRequest("missing data")
	}

	rates := a.GetRates(h)

	js, _ := json.Marshal(rates)
	return c.JSONBlob(http.StatusOK, js)
}
