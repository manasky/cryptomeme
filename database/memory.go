package database

import (
	"context"
	"cryptolist/common"
	"time"
)

type MemoryCache struct {
	markets []*common.Market
	time    time.Time
}

func init() {
	Register("memory", &MemoryCache{})
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
