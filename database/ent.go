package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/manasky/cryptomeme/common"
	"github.com/manasky/cryptomeme/internal/ent"
	"github.com/manasky/cryptomeme/internal/ent/marketchart"
	"github.com/manasky/cryptomeme/internal/ent/markets"
	"time"
)

type Ent struct {
	client *ent.Client
}

type EntConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func init() {
	Register("ent", &Ent{})
}

func (e *Ent) open(driver, username, password, host, port, database string) error {
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
		return err
	}

	e.client = ent.NewClient(ent.Driver(drv))
	return nil
}

func (e *Ent) Connect(ctx context.Context, conf interface{}) error {
	if c, ok := conf.(*EntConfig); ok {
		err := e.open(c.Driver, c.Username, c.Password, c.Host, c.Port, c.Database)
		if err != nil {
			return err
		}

		return e.Migrate(ctx)
	}
	return ErrInvalidDriverConfig
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

func (e *Ent) SaveMarketChart(ctx context.Context, chart *common.MarketChart) error {
	c, err := e.client.MarketChart.Query().Where(marketchart.Name(chart.Name)).Where(marketchart.Currency(chart.Currency)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_, err = e.client.MarketChart.Create().SetCurrency(chart.Currency).SetName(chart.Name).SetChart(*chart).Save(ctx)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	c.Chart = *chart
	_, err = c.Update().Save(ctx)
	return err
}

func (e *Ent) MarketsCharts(ctx context.Context, currency string) (map[string]*common.MarketChart, error) {
	charts, err := e.client.MarketChart.Query().Where(marketchart.Currency(currency)).All(ctx)
	if err != nil {
		return nil, err
	}

	tmp := make(map[string]*common.MarketChart)
	for _, c := range charts {
		tmp[c.Name] = &c.Chart
	}
	return tmp, nil
}
