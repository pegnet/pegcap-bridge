package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var names map[string]string

func init() {
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
	names["TWD"] = "Tiawanese Dollar"
	names["KRW"] = "Korean Won"
	names["ARS"] = "Argentine Peso"
	names["BRL"] = "Brazil Real"
	names["PHP"] = "Philippine Peso"
	names["MXN"] = "Mexican Peso"
	names["XAU"] = "Gold Troy Ounce"
	names["XAG"] = "Silver Troy Ounce"
	names["XPD"] = "Palladium Troy Ounce"
	names["XPT"] = "Platinum Troy Ounce"
	names["XBT"] = "Bitcoin"
	names["ETH"] = "Ethereum"
	names["LTC"] = "Litecoin"
	names["RVN"] = "Ravencoin"
	names["XBC"] = "Bitcoin Cash"
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
	"XAU",
	"XAG",
	"XPD",
	"XPT",
	"XBT",
	"ETH",
	"LTC",
	"RVN",
	"XBC",
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
	"XAU",
	"XAG",
	"XBT",
	"ETH",
	"LTC",
	"RVN",
	"XBC",
	"FCT",
	"BNB",
	"XLM",
	"ADA",
	"XMR",
	"DASH",
	"ZEC",
	"DCR",
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
