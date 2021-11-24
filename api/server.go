package api

import (
	"context"
	"encoding/json"
	"github.com/manasky/cryptomeme/database"
	"github.com/manasky/cryptomeme/meme"
	"github.com/manasky/cryptomeme/common"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	db database.Provider
	mm *meme.Manager
}

func NewServer(db database.Provider, memeManager *meme.Manager) *Server {
	return &Server{
		db: db,
		mm: memeManager,
	}
}

func (s *Server) Markets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	t := r.FormValue("time")
	var ti time.Time
	if t != "" {
		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			http.Error(w, "invalid time format", http.StatusBadRequest)
			return
		}

		ti = time.Unix(i, 0)
	}

	if ti.IsZero() {
		ti = time.Now()
	}
	markets, err := s.db.Markets(context.Background(), ti)
	if err != nil {
		log.Printf("error while fetching from database: %s", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	charts, err := s.db.MarketsCharts(context.Background(), "usd")
	if err != nil {
		log.Printf("error while fetching charts from database: %s", err)
	}

	for _, market := range markets {
		if charts != nil {
			if chart, ok := charts[market.ID]; ok {
				market.Prices7Days = chart.Prices()
			}
		}

		m, err := s.mm.Meme(strings.ToLower(market.Name), market.PriceChangePercentage24h)
		if err == nil {
			market.Meme = m.Image
			if c, err := meme.ParseCaption(m.Caption, market); err == nil {
				market.MemeCaption = c
			}

			market.MemeContent = []*common.Content{}
			if len(m.Content) > 0 {
				var memeContent []*common.Content
				for _, c := range m.Content {
					tmp := &common.Content{
						Text: c.Text,
						Image: c.Image,
						X: c.X,
						Y: c.Y,
					}
					if tmp.Text != "" {
						if txt, err := meme.ParseCaption(tmp.Text, market); err == nil {
							tmp.Text = txt
						}
					}
					memeContent = append(memeContent, tmp)
				}
				market.MemeContent = memeContent
			}
		}
	}

	resp, err := json.Marshal(markets)
	if err != nil {
		log.Printf("error while marshaling: %s", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Write(resp) //TODO
}
