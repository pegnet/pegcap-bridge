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
	names["pKRW"] = "Korean Won"
	names["pBRL"] = "Brazil Real"
	names["pPHP"] = "Philippine Peso"
	names["pMXN"] = "Mexican Peso"
	names["pGOLD"] = "Gold Troy Ounce"
	names["pSILVER"] = "Silver Troy Ounce"
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
	names["pAUD"] = "Australian Dollar"
	names["pNZD"] = "New Zealand Dollar"
	names["pSEK"] = "Swedish Krona"
	names["pNOK"] = "Norwegian Krone"
	names["pRUB"] = "Russian Rouble"
	names["pZAR"] = "South African Rand"
	names["pTRY"] = "Turkish Lira"
	names["pEOS"] = "EOS"
	names["pLINK"] = "Chainlink"
	names["pATOM"] = "Cosmos"
	names["pBAT"] = "Basic Attention Token"
	names["pXTZ"] = "Tezos"
	names["pHBAR"] = "Hedera Hashgraph"
	names["pNEO"] = "NEO"
	names["pCRO"] = "Crypto.com"
	names["pETC"] = "Ethereum Classic"
	names["pONT"] = "Ontology"
	names["pDOGE"] = "Dogecoin"
	names["pVET"] = "VeChain"
	names["pHT"] = "Huobi Token"
	names["pALGO"] = "Algorand"
	names["pDGB"] = "DigiByte"
	names["pAED"] = "United Arab Emirates Dirham"
	names["pARS"] = "Argentine Peso"
	names["pTWD"] = "Taiwanese Dollar"
	names["pRWF"] = "Rwandan Franc"
	names["pKES"] = "Kenyan Shilling"
	names["pUGX"] = "Ugandan Shilling"
	names["pTZS"] = "Tanzanian Shilling"
	names["pBIF"] = "Burundian Franc"
	names["pETB"] = "Ethiopian Birr"
	names["pNGN"] = "Nigerian Naira"
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
