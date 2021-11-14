package database

import (
	"context"
	"cryptolist/common"
	"cryptolist/internal/ent"
	"cryptolist/internal/ent/markets"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"time"
)

type Ent struct {
	client *ent.Client
}

func Open(driver, username, password, host, port, database string) (*Ent, error) {
	var dataSourceName string

	switch driver {
	case "sqlite", "sqlite3":
		dataSourceName = fmt.Sprintf("file:%s.sqlite3?_fk=1", database)
	case "postgres":
		dataSourceName = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			host,
			port,
			username,
			database,
			password,
		)
	default:
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			username,
			password,
			host,
			port,
			database,
		)
	}

	drv, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Ent{
		client: ent.NewClient(ent.Driver(drv)),
	}, nil
}

func (e *Ent) Migrate(ctx context.Context) error {
	return e.client.Schema.Create(ctx)
}

func (e *Ent) Close() error {
	return e.client.Close()
}

func (e *Ent) SaveMarkets(ctx context.Context, markets []*common.Market, t time.Time) error {
	_, err := e.client.Markets.Create().SetMarkets(markets).SetTime(t).Save(ctx)
	return err
}

func (e *Ent) Markets(ctx context.Context, t time.Time) ([]*common.Market, error) {
	mrks, err := e.client.Markets.Query().Where(markets.TimeLTE(t)).Order(ent.Desc(markets.FieldTime)).First(ctx)
	if err != nil {
		return nil, err
	}
	return mrks.Markets, nil
}
