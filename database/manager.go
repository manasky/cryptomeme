package database

import (
	"context"
	"cryptolist/common"
	"errors"
	"time"
)

var (
	ErrInvalidDriverConfig = errors.New("invalid driver configure file")
	ErrDriverNotFound      = errors.New("driver not found")

	drivers = make(map[string]Driver)
)

type Driver interface {
	Connect(context.Context, interface{}) error
	Provider
}

type Manager struct {
	drv Driver
}

func NewManager(ctx context.Context, driver string, conn interface{}) (*Manager, error) {
	if d, ok := drivers[driver]; ok {

		err := d.Connect(ctx, conn)
		if err != nil {
			return nil, err
		}
		return &Manager{
			drv: d,
		}, nil
	}
	return nil, ErrDriverNotFound
}

func (m *Manager) SaveMarkets(ctx context.Context, markets []*common.Market, t time.Time) error {
	return m.drv.SaveMarkets(ctx, markets, t)
}

func (m *Manager) Markets(ctx context.Context, t time.Time) ([]*common.Market, error) {
	return m.drv.Markets(ctx, t)
}

func Register(name string, driver Driver) {
	drivers[name] = driver
}
