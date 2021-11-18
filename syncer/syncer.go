package syncer

import (
	"context"
	"cryptolist/api"
	"cryptolist/database"
	"errors"
	"log"
	"time"
)

type Syncer struct {
	api api.API
	db  database.Provider

	stop   chan bool
	ticker *time.Ticker
	config *Config
}

type Config struct {
	Delay time.Duration
}

func New(api api.API, db database.Provider, config *Config) (*Syncer, error) {
	return &Syncer{
		api: api,
		db:  db,

		stop:   make(chan bool, 1),
		config: config,
	}, nil
}

func (m *Syncer) Start() error {
	if m.ticker != nil {
		return errors.New("already syncing")
	}

	m.ticker = time.NewTicker(m.config.Delay)

	go func() {
		m.syncAPI()
		for {
			select {
			case <-m.stop:
				return
			case t := <-m.ticker.C:
				log.Println("syncing at", t)
				m.syncAPI()
			}
		}
	}()
	return nil
}

func (m *Syncer) Stop() error {
	if m.ticker == nil {
		return errors.New("no syncer found")
	}
	m.stop <- true
	m.ticker.Stop()
	m.ticker = nil
	return nil
}

func (m *Syncer) syncAPI() {
	ctx := context.Background()
	markets, err := m.api.Markets(ctx, "usd", 10)
	if err != nil {
		log.Printf("error while fetching from API: %s", err)
	}

	err = m.db.SaveMarkets(ctx, markets, time.Now())
	if err != nil {
		log.Printf("error while saving to database: %s", err)
	}
}
