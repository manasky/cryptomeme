package api

import (
	"context"
	"github.com/manasky/cryptomeme/common"
)

type API interface {
	Markets(ctx context.Context, currency string, perPage int) ([]*common.Market, error)
	MarketChart(ctx context.Context, id, currency string, days uint) (*common.MarketChart, error)
}
