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
	translate["XAG"] = "Silver"
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
	names["pUSD"] = "US Dollar"
	names["pEUR"] = "Euro"
	names["pJPY"] = "Japanese Yen"
	names["pGBP"] = "Pound Sterling"
	names["pCAD"] = "Canadian Dollar"
	names["pCHF"] = "Swiss Franc"
	names["pINR"] = "Indian Rupee"
	names["pSGD"] = "Singapore Dollar"
	names["pCNY"] = "Chinese Yuan"
	names["pHKD"] = "Hong Kong Dollar"
	names["pTWD"] = "Taiwanese Dollar"
	names["pKRW"] = "Korean Won"
	names["pARS"] = "Argentine Peso"
	names["pBRL"] = "Brazil Real"
	names["pPHP"] = "Philippine Peso"
	names["pMXN"] = "Mexican Peso"
	names["pGOLD"] = "Gold Troy Ounce"
	names["pSILVER"] = "Silver Troy Ounce"
	names["pXPD"] = "Palladium Troy Ounce"
	names["pXPT"] = "Platinum Troy Ounce"
	names["pBTC"] = "Bitcoin"
	names["pETH"] = "Ethereum"
	names["pLTC"] = "Litecoin"
	names["pRVN"] = "Ravencoin"
	names["pBCH"] = "Bitcoin Cash"
	names["pFCT"] = "Factom"
	names["pBNB"] = "Binance Coin"
	names["pXLM"] = "Stellar"
	names["pADA"] = "Cardano"
	names["pXMR"] = "Monero"
	names["pDASH"] = "Dash"
	names["pZEC"] = "Zcash"
	names["pDCR"] = "Decred"
}

var V1Assets = []string{
	"PEG",
	"pUSD",
	"pEUR",
	"pJPY",
	"pGBP",
	"pCAD",
	"pCHF",
	"pINR",
	"pSGD",
	"pCNY",
	"pHKD",
	"pKRW",
	"pBRL",
	"pPHP",
	"pMXN",
	"pGOLD",
	"pSILVER",
	"pXPD",
	"pXPT",
	"pBTC",
	"pETH",
	"pLTC",
	"pRVN",
	"pBCH",
	"pFCT",
	"pBNB",
	"pXLM",
	"pADA",
	"pXMR",
	"pDASH",
	"pZEC",
	"pDCR",
}

var V2Assets = []string{
	"PEG",
	"pUSD",
	"pEUR",
	"pJPY",
	"pGBP",
	"pCAD",
	"pCHF",
	"pINR",
	"pSGD",
	"pCNY",
	"pHKD",
	"pKRW",
	"pBRL",
	"pPHP",
	"pMXN",
	"pGOLD",
	"pSILVER",
	"pBTC",
	"pETH",
	"pLTC",
	"pRVN",
	"pBCH",
	"pFCT",
	"pBNB",
	"pXLM",
	"pADA",
	"pXMR",
	"pDASH",
	"pZEC",
	"pDCR",
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
