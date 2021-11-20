package database

import (
	"context"
	"cryptolist/common"
	"errors"
	"time"
)

type MemoryCache struct {
	markets       []*common.Market
	marketsCharts map[string]map[string]*common.MarketChart
	time          time.Time
}

func init() {
	Register("memory", &MemoryCache{
		marketsCharts: make(map[string]map[string]*common.MarketChart),
	})
}

func (m *MemoryCache) Connect(ctx context.Context, c interface{}) error {
	return nil
}

func (m *MemoryCache) SaveMarkets(ctx context.Context, markets []*common.Market, t time.Time) error {
	m.markets = markets
	m.time = t
	return nil
}

func (m *MemoryCache) Markets(ctx context.Context, t time.Time) ([]*common.Market, error) {
	return m.markets, nil
}

func (m *MemoryCache) SaveMarketChart(ctx context.Context, chart *common.MarketChart) error {
	if m.marketsCharts[chart.Currency] == nil {
		m.marketsCharts[chart.Currency] = make(map[string]*common.MarketChart)
	}
	m.marketsCharts[chart.Currency][chart.Name] = chart
	return nil
}

func (m *MemoryCache) MarketsCharts(ctx context.Context, currency string) (map[string]*common.MarketChart, error) {
	if c, ok := m.marketsCharts[currency]; ok {
		return c, nil
	}
	return nil, errors.New("currency not found")
}
