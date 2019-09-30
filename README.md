# Dummy Bridge
Middleware for the link between Factomize and Pegnet. Generates dummy data.

## Data

The PegNet `Genesis` block is `206422`. Any height less than that is not support.

At block height `210330` the list of assets changes, removing XPD and XPT.

The following block heights have intentionally missing data: 212122, 206522, 210000, 212000, 212001, 212002, 212003

For the `rates` and `market` responses, the numbers are **64 bit floats** with precision of 8 decimals.


## REST Methods
### Errors

Errors return a 500 code along with a descriptive JSON error of `{"message": "<error message>"}`.

`:height` is an int between PegNet's genesis and the current block height. Other inputs will results in an error.

#### /v1/heights

Returns the current height along with that block's unix timestamp. Also includes the genesis block.

Example Response:
```json
{
    "Genesis":206422,
    "Current":212127,
    "Blocktime":1569766920
}
```

#### /v1/asset/list/:height

Returns an array of asset codes that are used at the given height.

Example Response:
```json
["PEG","USD","EUR","JPY","GBP","CAD","CHF","INR","SGD","CNY","HKD","KRW","BRL","PHP","MXN","XAU","XAG","XPD","XPT","XBT","ETH","LTC","RVN","XBC","FCT","BNB","XLM","ADA","XMR","DASH","ZEC","DCR"]
```

#### /v1/asset/names

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
    "XAG":"Silver Troy Ounce",
    "XAU":"Gold Troy Ounce",
    "XBC":"Bitcoin Cash",
    "XBT":"Bitcoin",
    "XLM":"Stellar",
    "XMR":"Monero",
    "XPD":"Palladium Troy Ounce",
    "XPT":"Platinum Troy Ounce",
    "ZEC":"Zcash"
}
```

#### /v1/asset/name/:code

Returns the full name of a specific asset code.

Example Response for `/v1/asset/name/USD`:
```json
{
    "USD":"US Dollar"
}
```

#### /v1/rates/:height

Returns the asset rates for a specific height. 

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
   "XAG":18.23120613,
   "XAU":1500.1900662,
   "XBC":298.59537544,
   "XBT":10501.35699081,
   "XLM":0.06424941,
   "XMR":78.52750335,
   "XPD":1550.28005521,
   "XPT":99.25,
   "ZEC":49.26302176
}
```

#### /v1/market/:height

Returns the market data for a specific height.

Example Response:
```json
{ 
   "Burnt":0,
   "Supply":{ 
      "ADA":48.74097526,
      "BNB":24429.08607167,
      "BRL":242.30168971,
      "CAD":746.3961165,
      "CHF":1013.38639359,
      "CNY":139.50255603,
      "DASH":87355.41594458,
      "DCR":24297.10120604,
      "ETH":179484.12400971,
      "EUR":1124.88683486,
      "FCT":5597.9255608,
      "GBP":1151.30978351,
      "HKD":126.77389305,
      "INR":13.85666571,
      "JPY":9.38041178,
      "KRW":0.77292582,
      "LTC":70201.62387039,
      "MXN":50.9544447,
      "PEG":0,
      "PHP":19.31926299,
      "RVN":35.38639879,
      "SGD":717.99538471,
      "USD":999.90996262,
      "XAG":18237.76135358,
      "XAU":1506207.75536337,
      "XBC":298810.35832777,
      "XBT":10570789.21202436,
      "XLM":63.88573348,
      "XMR":78394.76260323,
      "XPD":1558436.39016583,
      "XPT":452247.3024172,
      "ZEC":48931.11979955
   },
   "Volume":{ 
      "ADA":6.49374703,
      "BNB":1200.93588849,
      "BRL":23.27539549,
      "CAD":53.51700692,
      "CHF":115.55237429,
      "CNY":16.16434415,
      "DASH":6228.61865677,
      "DCR":1169.15483812,
      "ETH":18049.53973447,
      "EUR":5.09686895,
      "FCT":667.6715148,
      "GBP":11.39345225,
      "HKD":7.28431414,
      "INR":0.85830606,
      "JPY":0.95881592,
      "KRW":0.04510752,
      "LTC":5131.37710198,
      "MXN":6.4138677,
      "PEG":0,
      "PHP":1.37251855,
      "RVN":3.18145955,
      "SGD":77.7360612,
      "USD":5.17812219,
      "XAG":27.78456726,
      "XAU":77494.68045714,
      "XBC":4155.98562541,
      "XBT":1124740.16552694,
      "XLM":9.48679649,
      "XMR":5003.97543396,
      "XPD":141092.32803746,
      "XPT":18863.71474459,
      "ZEC":6481.41377218
   }
}
```