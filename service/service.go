package service

import (
	"context"
	"github.com/manasky/cryptomeme/api"
	"github.com/manasky/cryptomeme/database"
	"log"
	"time"
)

type Service struct {
	api api.API
	db  database.Provider

	retryDelay time.Duration
}

func NewService(api api.API, db database.Provider) *Service {
	return &Service{
		api:        api,
		db:         db,
		retryDelay: 10 * time.Second,
	}
}

func (s *Service) SyncMarkets() {
	ctx := context.Background()
	markets, err := s.api.Markets(ctx, "usd", 20)
	if err != nil {
		log.Printf("error while fetching from API: %s", err)
		s.retry(s.SyncMarkets)
		return
	}

	err = s.db.SaveMarkets(ctx, markets, time.Now())
	if err != nil {
		log.Printf("error while saving to database: %s", err)
		s.retry(s.SyncMarkets)
		return
	}
}

func (s *Service) SyncMarketsChart() {
	ctx := context.Background()
	markets, err := s.db.Markets(ctx, time.Now())
	if err != nil {
		log.Printf("error while fetching markets from the database: %s, retry in %s", err, s.retryDelay)
		s.retry(s.SyncMarketsChart)
		return
	}

	if markets == nil {
		log.Printf("error while fetching markets from the database: no record found, retry in %s", s.retryDelay)
		s.retry(s.SyncMarketsChart)
	}

	for _, market := range markets {
		chart, err := s.api.MarketChart(ctx, market.ID, "usd", 7)
		if err != nil {
			log.Printf("error while fetching %s market chart from the API: %s", market.ID, err)
			continue
		}

		go func() {
			err = s.db.SaveMarketChart(ctx, chart)
			if err != nil {
				log.Printf("error while saving market charts to database: %s", err)
				return
			}
		}()

		time.Sleep(200 * time.Millisecond) // sleep to not interrupt the API
	}
}

func (s *Service) retry(f func()) {
	time.AfterFunc(s.retryDelay, f)
}
