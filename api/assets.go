package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

var translate map[string]string
var rtranslate map[string]string
var names map[string]string

func init() {
	translate = make(map[string]string)
	translate["pXAU"] = "pGOLD"
	translate["pXAG"] = "pSILVER"
	translate["pXBC"] = "pBCH"
	translate["pXBT"] = "pBTC"

	rtranslate = make(map[string]string)
	rtranslate["pGOLD"] = "pXAU"
	rtranslate["pSILVER"] = "pXAG"
	rtranslate["pBCH"] = "pXBC"
	rtranslate["pBTC"] = "pXBT"

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

func (api *Api) AssetNames(c echo.Context) error {
	js, _ := json.Marshal(names)
	return c.JSONBlob(http.StatusOK, js)
}

func Trans(key string) string {
	if n, ok := translate[key]; ok {
		return n
	}
	return key
}

func RTrans(key string) string {
	if n, ok := rtranslate[key]; ok {
		return n
	}
	return key
}
