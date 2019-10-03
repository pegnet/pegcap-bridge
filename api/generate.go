package api

import (
	"math/rand"
	"time"
)

type Bound struct {
	Code   string
	Min    uint64
	Max    uint64
	Factor uint64
}

var Bounds = []Bound{
	Bound{
		Code:   "PEG",
		Min:    0,
		Max:    0,
		Factor: 0,
	},
	Bound{
		Code:   "pUSD",
		Min:    100000000,
		Max:    100000000,
		Factor: 0,
	},
	Bound{
		Code:   "pEUR",
		Min:    105000000,
		Max:    120000000,
		Factor: 17000,
	},
	Bound{
		Code:   "pJPY",
		Min:    900000,
		Max:    980000,
		Factor: 175,
	},
	Bound{
		Code:   "pGBP",
		Min:    100000000,
		Max:    130000000,
		Factor: 30000,
	},
	Bound{
		Code:   "pCAD",
		Min:    74000000,
		Max:    76000000,
		Factor: 10000,
	},
	Bound{
		Code:   "pCHF",
		Min:    100000000,
		Max:    103000000,
		Factor: 15000,
	},
	Bound{
		Code:   "pINR",
		Min:    1300000,
		Max:    1450000,
		Factor: 500,
	},
	Bound{
		Code:   "pSGD",
		Min:    71000000,
		Max:    73000000,
		Factor: 7000,
	},
	Bound{
		Code:   "pCNY",
		Min:    13500000,
		Max:    14200000,
		Factor: 700,
	},
	Bound{
		Code:   "pHKD",
		Min:    12700000,
		Max:    12800000,
		Factor: 350,
	},
	Bound{
		Code:   "pKRW",
		Min:    75000,
		Max:    80000,
		Factor: 50,
	},
	Bound{
		Code:   "pBRL",
		Min:    23000000,
		Max:    25000000,
		Factor: 5500,
	},
	Bound{
		Code:   "pPHP",
		Min:    1900000,
		Max:    1940000,
		Factor: 400,
	},
	Bound{
		Code:   "pMXN",
		Min:    4900000,
		Max:    5200000,
		Factor: 1100,
	},
	Bound{
		Code:   "pGOLD",
		Min:    145000000000,
		Max:    155000000000,
		Factor: 55000000,
	},
	Bound{
		Code:   "pSILVER",
		Min:    1686340000,
		Max:    1961590000,
		Factor: 1276683,
	},
	Bound{
		Code:   "pXPD",
		Min:    145000000000,
		Max:    165000000000,
		Factor: 34904736,
	},
	Bound{
		Code:   "pXPT",
		Min:    80000000000,
		Max:    10000000000,
		Factor: 25000000,
	},
	Bound{
		Code:   "pBTC",
		Min:    900000000000,
		Max:    1200000000000,
		Factor: 900000000,
	},
	Bound{
		Code:   "pETH",
		Min:    16000000000,
		Max:    20000000000,
		Factor: 21892737,
	},
	Bound{
		Code:   "pLTC",
		Min:    6268160000,
		Max:    7843500000,
		Factor: 10418599,
	},
	Bound{
		Code:   "pRVN",
		Min:    2990000,
		Max:    4140000,
		Factor: 10796,
	},
	Bound{
		Code:   "pBCH",
		Min:    27387500000,
		Max:    32267500000,
		Factor: 39785821,
	},
	Bound{
		Code:   "pFCT",
		Min:    295260000,
		Max:    837120000,
		Factor: 1324939,
	},
	Bound{
		Code:   "pBNB",
		Min:    1996180000,
		Max:    2882880000,
		Factor: 3487532,
	},
	Bound{
		Code:   "pXLM",
		Min:    5530000,
		Max:    7310000,
		Factor: 21219,
	},
	Bound{
		Code:   "pADA",
		Min:    4340000,
		Max:    5260000,
		Factor: 7783,
	},
	Bound{
		Code:   "pXMR",
		Min:    6629530000,
		Max:    9062430000,
		Factor: 11197990,
	},
	Bound{
		Code:   "pDASH",
		Min:    7864730000,
		Max:    9669730000,
		Factor: 13370280,
	},
	Bound{
		Code:   "pZEC",
		Min:    4392930000,
		Max:    5457300000,
		Factor: 9116174,
	},
	Bound{
		Code:   "pDCR",
		Min:    2183270000,
		Max:    2682370000,
		Factor: 3936380,
	},
}

var BoundMap map[string]Bound
var rng *rand.Rand

func init() {
	BoundMap = make(map[string]Bound)
	for _, b := range Bounds {
		BoundMap[b.Code] = b
	}
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func startRates() []TokenRate {
	var r []TokenRate
	for _, b := range Bounds {
		r = append(r, TokenRate{
			Name: b.Code,
			Rate: (b.Min + b.Max) / 2,
		})
	}
	return r
}

func startMarket() Market {
	var m Market
	m.Burnt = 100e8
	for _, b := range Bounds {
		m.Info = append(m.Info, MarketData{
			Name:   b.Code,
			Supply: (b.Min + b.Max) * 500,
			Volume: 0,
		})
	}
	return m
}

func (a *Api) Generate(height int) {
	if height < Genesis {
		return
	}

	md := a.GetMetadata()
	if md.Last >= height {
		return
	}

	if height == Genesis {
		a.SetMarket(Genesis, startMarket())
		a.SetRates(Genesis, startRates())
		md.Last = Genesis
		a.SetMetadata(md)
		return
	}

	a.Generate(height - 1)

	m := a.GetMarket(height - 1)
	if rng.Float64() < .25 {
		m.Burnt = 200e8 + uint64(rng.Int63n(250e8))
	} else {
		m.Burnt = 0
	}

	for i, d := range m.Info {
		d.Volume = uint64(rng.Float64() * .15 * float64(d.Supply))
		change := (rng.Float64() - .48) / 4 * float64(d.Volume)
		if change < 0 && -change >= float64(d.Supply) {
			change = -change
		}

		d.Supply += uint64(change)

		m.Info[i] = d
	}
	a.SetMarket(height, m)

	prev := a.GetRates(height - 1)
	for j, r := range prev {
		bound := BoundMap[r.Name]
		if bound.Factor == 0 {
			continue
		}
		step := uint64(rng.Int63n(int64(bound.Factor)))
		val := r.Rate

		if rng.Int31n(2) == 1 {
			val += step
		} else {
			val -= step
		}

		if val <= bound.Min {
			val = bound.Min + 3*bound.Factor
		}

		if val >= bound.Max {
			val = bound.Max - 3*bound.Factor
		}

		r.Rate = val
		prev[j] = r
	}

	a.SetRates(height, prev)

	md.Last = height
	a.SetMetadata(md)
}
