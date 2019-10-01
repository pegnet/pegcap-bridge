package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var translate map[string]string
var names map[string]string
var V1Exists map[string]bool
var V2Exists map[string]bool

func init() {
	translate = make(map[string]string)
	translate["XAU"] = "GOLD"
	translate["XAG"] = "SILVER"
	translate["XBC"] = "BCH"
	translate["XBT"] = "BTC"

	V1Exists = make(map[string]bool)
	for _, a := range V1Assets {
		V1Exists[a] = true
	}
	V2Exists = make(map[string]bool)
	for _, a := range V2Assets {
		V2Exists[a] = true
	}
	names = make(map[string]string)
	names["PEG"] = "PegNet"
	names["USD"] = "US Dollar"
	names["EUR"] = "Euro"
	names["JPY"] = "Japanese Yen"
	names["GBP"] = "Pound Sterling"
	names["CAD"] = "Canadian Dollar"
	names["CHF"] = "Swiss Franc"
	names["INR"] = "Indian Rupee"
	names["SGD"] = "Singapore Dollar"
	names["CNY"] = "Chinese Yuan"
	names["HKD"] = "Hong Kong Dollar"
	names["TWD"] = "Taiwanese Dollar"
	names["KRW"] = "Korean Won"
	names["ARS"] = "Argentine Peso"
	names["BRL"] = "Brazil Real"
	names["PHP"] = "Philippine Peso"
	names["MXN"] = "Mexican Peso"
	names["GOLD"] = "Gold Troy Ounce"
	names["SILVER"] = "Silver Troy Ounce"
	names["XPD"] = "Palladium Troy Ounce"
	names["XPT"] = "Platinum Troy Ounce"
	names["BTC"] = "Bitcoin"
	names["ETH"] = "Ethereum"
	names["LTC"] = "Litecoin"
	names["RVN"] = "Ravencoin"
	names["BCH"] = "Bitcoin Cash"
	names["FCT"] = "Factom"
	names["BNB"] = "Binance Coin"
	names["XLM"] = "Stellar"
	names["ADA"] = "Cardano"
	names["XMR"] = "Monero"
	names["DASH"] = "Dash"
	names["ZEC"] = "Zcash"
	names["DCR"] = "Decred"
}

var V1Assets = []string{
	"PEG",
	"USD",
	"EUR",
	"JPY",
	"GBP",
	"CAD",
	"CHF",
	"INR",
	"SGD",
	"CNY",
	"HKD",
	"KRW",
	"BRL",
	"PHP",
	"MXN",
	"GOLD",
	"SILVER",
	"XPD",
	"XPT",
	"BTC",
	"ETH",
	"LTC",
	"RVN",
	"BCH",
	"FCT",
	"BNB",
	"XLM",
	"ADA",
	"XMR",
	"DASH",
	"ZEC",
	"DCR",
}

var V2Assets = []string{
	"PEG",
	"USD",
	"EUR",
	"JPY",
	"GBP",
	"CAD",
	"CHF",
	"INR",
	"SGD",
	"CNY",
	"HKD",
	"KRW",
	"BRL",
	"PHP",
	"MXN",
	"GOLD",
	"SILVER",
	"BTC",
	"ETH",
	"LTC",
	"RVN",
	"BCH",
	"FCT",
	"BNB",
	"XLM",
	"ADA",
	"XMR",
	"DASH",
	"ZEC",
	"DCR",
}

func (api *Api) AssetExists(height int, asset string) bool {
	if height >= V2 {
		return V2Exists[asset]
	}
	return V1Exists[asset]
}

func (api *Api) AssetNames(c echo.Context) error {
	js, _ := json.Marshal(names)
	return c.JSONBlob(http.StatusOK, js)
}

func (api *Api) AssetList(c echo.Context) error {
	raw := c.Param("height")
	height, err := strconv.Atoi(raw)
	if err != nil {
		return api.BadRequest("invalid height")
	}

	if height >= 210330 {
		js, _ := json.Marshal(V2Assets)
		return c.JSONBlob(http.StatusOK, js)
	}

	js, _ := json.Marshal(V1Assets)
	return c.JSONBlob(http.StatusOK, js)
}

func (api *Api) AssetName(c echo.Context) error {
	name, ok := names[c.Param("code")]
	if !ok {
		return api.BadRequest("invalid code")
	}
	oneoff := make(map[string]string)
	oneoff[c.Param("code")] = name
	js, _ := json.Marshal(oneoff)
	return c.JSONBlob(http.StatusOK, js)
}
