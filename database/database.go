package database

import (
	"context"
	"cryptolist/common"
	"time"
)

type Provider interface {
	SaveMarkets(ctx context.Context, markets []*common.Market, t time.Time) error
	Markets(ctx context.Context, t time.Time) ([]*common.Market, error)
}
