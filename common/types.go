package common

import "time"

type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type Market struct {
	ID                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 float64   `json:"current_price"`
	MarketCap                    float64   `json:"market_cap"`
	MarketCapRank                uint64    `json:"market_cap_rank"`
	FullyDilutedValuation        float64   `json:"fully_diluted_valuation"`
	TotalVolume                  float64   `json:"total_volume"`
	High24h                      float64   `json:"high_24h"`
	Low24h                       float64   `json:"low_24h"`
	PriceChange24h               float64   `json:"price_change_24h"`
	MarketCapChange24h           float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64   `json:"price_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  float64   `json:"total_supply"`
	MaxSupply                    float64   `json:"max_supply"`
	ATH                          float64   `json:"ath"`
	ATHChangePercentage          float64   `json:"ath_change_percentage"`
	ATHDate                      time.Time `json:"ath_date"`
	ATL                          float64   `json:"atl"`
	ATLChangePercentage          float64   `json:"atl_change_percentage"`
	ATLDate                      time.Time `json:"atl_date"`
	ROI                          *ROI      `json:"roi"`
	LastUpdated                  time.Time `json:"last_updated"`
	
	Meme string `json:"meme"`
	MemeCaption string `json:"meme_caption"`
	Prices7Days	[]*ChartPrice `json:"prices_7d"`
}

type ChartPrice struct {
	Time int64 `json:"time"`
	Price float64 `json:"price"`
}

type MarketChart struct {
	Name string `json:"name"`
	Currency string `json:"currency"`
	RawPrices [][]float64       `json:"prices"`
	RawMarketCaps [][]float64   `json:"market_caps"`
	RawTotalVolumes [][]float64 `json:"total_volumes"`
}

func (m *MarketChart) Prices() []*ChartPrice {
	var tmp []*ChartPrice
	for _, p := range m.RawPrices {
		tmp = append(tmp, &ChartPrice{
			Time: int64(p[0]),
			Price: p[1],
		})
	}
	return tmp
}

func (m *MarketChart) MarketCaps() []*ChartPrice {
	var tmp []*ChartPrice
	for _, p := range m.RawMarketCaps {
		tmp = append(tmp, &ChartPrice{
			Time: int64(p[0]),
			Price: p[1],
		})
	}
	return tmp
}

func (m *MarketChart) TotalVolumes() []*ChartPrice {
	var tmp []*ChartPrice
	for _, p := range m.RawTotalVolumes {
		tmp = append(tmp, &ChartPrice{
			Time: int64(p[0]),
			Price: p[1],
		})
	}
	return tmp
}



