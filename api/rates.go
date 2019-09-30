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

	rawrates := a.GetRates(h)

	rates := make(map[string]float64)
	for _, r := range rawrates {
		if a.AssetExists(h, r.Name) {
			rates[r.Name] = Uint64ToFloat(r.Rate)
		}
	}

	js, _ := json.Marshal(rates)
	return c.JSONBlob(http.StatusOK, js)
}
