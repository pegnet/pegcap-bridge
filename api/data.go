package api

import (
	"encoding/json"
	"fmt"
)

type Metadata struct {
	Last int
}

type TokenRate struct {
	Name string
	Rate uint64
}

type Market struct {
	Burnt uint64
	Info  []MarketData
}

type MarketData struct {
	Name   string
	Supply uint64
	Volume uint64
}

func SKey(bits ...string) []byte {
	if len(bits) == 0 {
		return []byte{}
	}

	r := []byte(bits[0])
	for i := 1; i < len(bits); i++ {
		r = append(r, '-')
		r = append(r, []byte(bits[i])...)
	}
	return r
}

func (a *Api) GenericGetHeight(v interface{}, height int, bits ...string) error {
	bits = append([]string{fmt.Sprintf("%d", height)}, bits...)
	return a.GenericGet(v, bits...)
}
func (a *Api) GenericGet(v interface{}, bits ...string) error {
	val, err := a.DB.Get(SKey(bits...), nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(val, v)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) GenericSet(v interface{}, bits ...string) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	err = a.DB.Put(SKey(bits...), data, nil)
	if err != nil {
		panic("unable to write to db: " + string(SKey(bits...)) + "; " + err.Error())
	}
}

func (a *Api) GenericSetHeight(v interface{}, height int, bits ...string) {
	bits = append([]string{fmt.Sprintf("%d", height)}, bits...)
	a.GenericSet(v, bits...)
}

func (a *Api) GetMetadata() Metadata {
	var m Metadata
	err := a.GenericGet(&m, "metadata")
	if err != nil {
		a.SetMetadata(Metadata{})
		return a.GetMetadata()
	}
	return m
}

func (a *Api) SetMetadata(m Metadata) {
	a.GenericSet(m, "metadata")
}

func (a *Api) GetRates(height int) []TokenRate {
	var rates []TokenRate
	a.GenericGetHeight(&rates, height, "rates")
	return rates
}

func (a *Api) SetRates(height int, rates []TokenRate) {
	a.GenericSetHeight(rates, height, "rates")
}

func (a *Api) GetMarket(height int) Market {
	var market Market
	a.GenericGetHeight(&market, height, "market")
	return market
}

func (a *Api) SetMarket(height int, market Market) {
	a.GenericSetHeight(&market, height, "market")
}
