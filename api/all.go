package api

import (
	"encoding/json"
	"net/http"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/labstack/echo/v4"
	"github.com/pegnet/pegnetd/fat/fat2"
	"github.com/pegnet/pegnetd/node/pegnet"
	"github.com/pegnet/pegnetd/srv"
)

const (
	VOLUME int = iota
	VOLUMEIN
	VOLUMEOUT
	VOLUMETX
	SUPPLY
	PRICE
)

type All struct {
	Height           int
	Blocktime        int64
	Burnt            float64
	TotalConversions float64
	Data             map[string][6]float64
}

func (a *Api) _getStats(height uint32) (*pegnet.Stats, error) {
	params := &srv.ParamsGetStats{Height: &height}
	var res pegnet.Stats
	err := a.Cli.Request("get-stats", params, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil

}

func (a *Api) _getRates(height uint32) (*srv.ResultPegnetTickerMap, error) {
	var res srv.ResultPegnetTickerMap
	err := a.Cli.Request("get-pegnet-rates", srv.ParamsGetPegnetRates{Height: &height}, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *Api) _blockTime(height uint32) (int64, error) {
	var dblock factom.DBlock
	dblock.Height = height
	if err := dblock.Get(a.Factom); err != nil {
		return 0, a.BadRequest("unable to contact factomd endpoint: " + err.Error())
	}

	return dblock.Timestamp.Unix(), nil
}

func (a *Api) All(c echo.Context) error {
	h, err := a.VerifyHeight(c)
	if err != nil {
		return err
	}

	stats, err := a._getStats(uint32(h))
	if err != nil {
		return a.BadRequest("unable to retrieve stats: " + err.Error())
	}

	rates, err := a._getRates(uint32(h))
	if err != nil {
		rates = nil
		if err.Error() != "jsonrpc2.Error{Code:-32809, Message:\"Not Found\", Data:\"could not find what you were looking for\"}" {
			return a.BadRequest("unable to retrieve rates: " + err.Error())
		}
	}

	bt, err := a._blockTime(uint32(h))
	if err != nil {
		return err
	}

	var all All
	all.Height = h
	all.Burnt = Uint64ToFloat(stats.Burns)
	all.Blocktime = bt
	all.Data = make(map[string][6]float64)

	if rates != nil {
		for k, v := range map[fat2.PTicker]uint64(*rates) {
			key := Trans(k.String())
			m := all.Data[key]
			m[PRICE] = Uint64ToFloat(v)
			all.Data[key] = m
		}
	}

	for k, v := range stats.Volume {
		k = Trans(k)
		m := all.Data[k]
		m[VOLUME] = Uint64ToFloat(v)
		all.Data[k] = m
	}

	for k, v := range stats.VolumeIn {
		k = Trans(k)
		m := all.Data[k]
		m[VOLUMEIN] = Uint64ToFloat(v)
		all.Data[k] = m
	}
	for k, v := range stats.VolumeOut {
		k = Trans(k)
		m := all.Data[k]
		m[VOLUMEOUT] = Uint64ToFloat(v)
		all.Data[k] = m

		all.TotalConversions += m[VOLUMEOUT] * m[PRICE]
	}

	for k, v := range stats.VolumeTx {
		k = Trans(k)
		m := all.Data[k]
		m[VOLUMETX] = Uint64ToFloat(v)
		all.Data[k] = m
	}

	for k, v := range stats.Supply {
		k = Trans(k)
		m := all.Data[k]
		m[SUPPLY] = Uint64ToFloat(uint64(v))
		all.Data[k] = m
	}

	js, _ := json.Marshal(all)
	return c.JSONBlob(http.StatusOK, js)
}
