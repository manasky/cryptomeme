package api

import (
	"context"
	"cryptolist/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	db database.Provider
}

func NewServer(db database.Provider) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Markets(w http.ResponseWriter, r *http.Request) {
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

	resp, err := json.Marshal(markets)
	if err != nil {
		log.Printf("error while marshaling: %s", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Write(resp) //TODO
}
