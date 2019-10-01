# Dummy Bridge
Middleware for the link between Factomize and Pegnet. Generates dummy data.

## Data

The PegNet `Genesis` block is `206422`. Any height less than that is not support.

At block height `210330` the list of assets changes, removing XPD and XPT.

The following block heights have intentionally missing data: 212122, 206522, 210000, 212000, 212001, 212002, 212003. These will result in status code 500 responses with the message: `{"message":"missing data"}`.

For the `rates` and `market` responses, the numbers are **64 bit floats** with precision of 8 decimals.


# REST Methods
## Errors

Errors return a 500 code along with a descriptive JSON error of `{"message": "<error message>"}`.

`:height` is an int between PegNet's genesis and the current block height. Other inputs will results in an error.

## /v1/heights

Returns the current height along with that block's unix timestamp. Also includes the genesis block.

Example Response:
```json
{
    "Genesis":206422,
    "Current":212127,
    "Blocktime":1569766920
}
```

## /v1/asset/list/:height

Returns an array of asset codes that are used at the given height.

Example Response:
```json
["PEG","USD","EUR","JPY","GBP","CAD","CHF","INR","SGD","CNY","HKD","KRW","BRL","PHP","MXN","GOLD","SILVER","XPD","XPT","BTC","ETH","LTC","RVN","BCH","FCT","BNB","XLM","ADA","XMR","DASH","ZEC","DCR"]
```

## /v1/asset/names

Returns the full name of all known asset codes.

Example Response:
```json
{ 
    "ADA":"Cardano",
    "ARS":"Argentine Peso",
    "BNB":"Binance Coin",
    "BRL":"Brazil Real",
    "CAD":"Canadian Dollar",
    "CHF":"Swiss Franc",
    "CNY":"Chinese Yuan",
    "DASH":"Dash",
    "DCR":"Decred",
    "ETH":"Ethereum",
    "EUR":"Euro",
    "FCT":"Factom",
    "GBP":"Pound Sterling",
    "HKD":"Hong Kong Dollar",
    "INR":"Indian Rupee",
    "JPY":"Japanese Yen",
    "KRW":"Korean Won",
    "LTC":"Litecoin",
    "MXN":"Mexican Peso",
    "PEG":"PegNet",
    "PHP":"Philippine Peso",
    "RVN":"Ravencoin",
    "SGD":"Singapore Dollar",
    "TWD":"Taiwanese Dollar",
    "USD":"US Dollar",
    "SILVER":"Silver Troy Ounce",
    "GOLD":"Gold Troy Ounce",
    "BCH":"Bitcoin Cash",
    "BTC":"Bitcoin",
    "XLM":"Stellar",
    "XMR":"Monero",
    "XPD":"Palladium Troy Ounce",
    "XPT":"Platinum Troy Ounce",
    "ZEC":"Zcash"
}
```

## /v1/asset/name/:code

Returns the full name of a specific asset code.

Example Response for `/v1/asset/name/USD`:
```json
{
    "USD":"US Dollar"
}
```

## /v1/rates/:height

Returns the asset rates for a specific height. All rates use USD as base. An FCT rate of `5.6515013` 
means that `1 FCT = 5.6515013 USD`.

Example Response:
```json
{ 
   "ADA":0.04794587,
   "BNB":24.37402664,
   "BRL":0.23998317,
   "CAD":0.74991207,
   "CHF":1.01502128,
   "CNY":0.13850009,
   "DASH":87.60382102,
   "DCR":24.36121272,
   "ETH":179.86516646,
   "EUR":1.12485646,
   "FCT":5.6515013,
   "GBP":1.14982772,
   "HKD":0.12749848,
   "INR":0.01374897,
   "JPY":0.00939842,
   "KRW":0.00077527,
   "LTC":70.61326566,
   "MXN":0.05049019,
   "PEG":0,
   "PHP":0.0192035,
   "RVN":0.0356818,
   "SGD":0.72001767,
   "USD":1,
   "SILVER":18.23120613,
   "GOLD":1500.1900662,
   "BCH":298.59537544,
   "BTC":10501.35699081,
   "XLM":0.06424941,
   "XMR":78.52750335,
   "XPD":1550.28005521,
   "XPT":99.25,
   "ZEC":49.26302176
}
```

## /v1/market/:height

Returns the market data for a specific height. Burnt is the additional amount of pFCT created that block.

Example Response:
```json
{ 
   "Burnt":345.18424661,
   "Supply":{ 
      "ADA":118.28260508,
      "BNB":56904.86664787,
      "BRL":301.83408013,
      "CAD":1095.89729937,
      "CHF":1385.46577659,
      "CNY":194.13177913,
      "DASH":156673.67219147,
      "DCR":26092.0461741,
      "ETH":258113.96044042,
      "EUR":1529.07384548,
      "FCT":11909.55299106,
      "GBP":2262.28967276,
      "HKD":180.89861775,
      "INR":30.46748947,
      "JPY":12.91986324,
      "KRW":1.73511295,
      "LTC":115817.67512238,
      "MXN":63.19019873,
      "PEG":0,
      "PHP":41.72209138,
      "RVN":58.54796461,
      "SGD":1026.22660097,
      "USD":1430.52875042,
      "SILVER":23816.55580375,
      "GOLD":1766646.41049297,
      "BCH":482377.36300028,
      "BTC":22661673.33739648,
      "XLM":106.95788701,
      "XMR":105565.66219623,
      "XPD":1945850.07114154,
      "XPT":973352.21987128,
      "ZEC":105276.11512017
   },
   "Volume":{ 
      "ADA":16.18870466,
      "BNB":5711.5387908,
      "BRL":35.04963485,
      "CAD":60.49312431,
      "CHF":51.64007972,
      "CNY":22.38208789,
      "DASH":23184.91790606,
      "DCR":1857.38150916,
      "ETH":36867.88860246,
      "EUR":176.04728955,
      "FCT":270.981585,
      "GBP":103.72879128,
      "HKD":20.16194456,
      "INR":2.78546782,
      "JPY":1.52047854,
      "KRW":0.21358942,
      "LTC":8334.63895588,
      "MXN":8.14553136,
      "PEG":0,
      "PHP":3.87022583,
      "RVN":2.88722763,
      "SGD":93.43208389,
      "USD":71.21949977,
      "SILVER":94.21417265,
      "GOLD":64182.6422265,
      "BCH":70168.58599648,
      "BTC":1308154.90972472,
      "XLM":6.10595405,
      "XMR":10448.66820555,
      "XPD":112364.02807111,
      "XPT":138012.26169199,
      "ZEC":14430.18929303
   }
}
```

## /v1/24hour/:height

Returns the approximate height 24 hours ago. Returns an error if 24 hours is earlier than the Genesis. 

Example Response for `/v1/24hour/207766`:
```json
207622
```


## /v1/all/:height

Returns a combination of height, market, and rates. 

Example Response:
```json
{ 
   "Blocktime":1568652300,
   "Burnt":345.18424661,
   "Supply":{ 
      // ...
   },
   "Volume":{ 
      // ...
   },
   "Rates":{ 
      // ...
   }
}
```