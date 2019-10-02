package api

import (
	"encoding/json"
	"net/http"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/labstack/echo/v4"
)

type All struct {
	Height    int
	Blocktime int
	Burnt     float64
	Supply    map[string]float64
	Volume    map[string]float64
	Rates     map[string]float64
}

func (a *Api) All(c echo.Context) error {
	h, err := a.VerifyHeight(c)
	if err != nil {
		return err
	}
	a.Generate(h)
	if Outage[h] {
		return a.BadRequest("missing data")
	}

	var all All

	all.Height = h

	rawrates := a.GetRates(h)
	all.Rates = make(map[string]float64)
	for _, r := range rawrates {
		if a.AssetExists(h, r.Name) {
			all.Rates[r.Name] = Uint64ToFloat(r.Rate)
		}
	}

	rawmarket := a.GetMarket(h)

	all.Supply = make(map[string]float64)
	all.Volume = make(map[string]float64)
	all.Burnt = Uint64ToFloat(rawmarket.Burnt)
	for _, md := range rawmarket.Info {
		if a.AssetExists(h, md.Name) {
			all.Supply[md.Name] = Uint64ToFloat(md.Supply)
			all.Volume[md.Name] = Uint64ToFloat(md.Volume)
		}
	}

	var dblock factom.DBlock
	dblock.Header.Height = uint32(h)
	dblock.Get(a.C)

	if !dblock.IsPopulated() {
		return a.BadRequest("unable to contact endpoint")
	}

	all.Blocktime = int(dblock.Header.Timestamp.Unix())

	js, _ := json.Marshal(all)
	return c.JSONBlob(http.StatusOK, js)
}
