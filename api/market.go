package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MarketResponse struct {
	Burnt  uint64
	Supply map[string]float64
	Volume map[string]float64
}

func (a *Api) Market(c echo.Context) error {
	h, err := a.VerifyHeight(c)
	if err != nil {
		return err
	}

	a.Generate(h)
	if Outage[h] {
		return a.BadRequest("missing data")
	}

	rawmarket := a.GetMarket(h)

	var market MarketResponse
	market.Supply = make(map[string]float64)
	market.Volume = make(map[string]float64)
	market.Burnt = rawmarket.Burnt
	for _, md := range rawmarket.Info {
		if a.AssetExists(h, md.Name) {
			market.Supply[md.Name] = Uint64ToFloat(md.Supply)
			market.Volume[md.Name] = Uint64ToFloat(md.Volume)
		}
	}

	js, _ := json.Marshal(market)
	return c.JSONBlob(http.StatusOK, js)
}
