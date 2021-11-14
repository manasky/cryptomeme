package api

import (
	"context"
	"cryptolist/common"
)

type API interface {
	Markets(ctx context.Context, currency string, perPage int) ([]*common.Market, error)
}
