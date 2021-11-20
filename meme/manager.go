package meme

import (
	"cryptolist/common"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type rangeCache struct {
	Range     float64
	Index     int
	ExpiredAt time.Time
}

type Manager struct {
	ds        *common.MemeDataset
	mp        map[string]int
	randCache map[string]*rangeCache
	cd        time.Duration // cache duration
}

func New(dataset string, cacheDuration time.Duration) (*Manager, error) {
	ds, err := readDataset(dataset)
	if err != nil {
		return nil, err
	}

	return &Manager{
		ds:        ds,
		mp:        indexDataset(ds),
		randCache: make(map[string]*rangeCache),
		cd:        cacheDuration,
	}, nil
}

func (m *Manager) Meme(coin string, change float64) (*common.Meme, error) {
	if c, ok := m.mp[coin]; ok {
		for _, r := range m.ds.Coins[c].Ranges {
			if change <= r.Value {

				return r.Memes[m.randOrCache(coin, r.Value, len(r.Memes))], nil
			}
		}
	}
	return &common.Meme{}, nil //todo
}

func (m *Manager) randOrCache(coin string, rng float64, len int) int {
	if c, ok := m.randCache[coin]; ok {
		if c.Range == rng && time.Now().Before(c.ExpiredAt) {
			return c.Index
		}
	}

	rand.Seed(time.Now().UnixNano())
	tmp := &rangeCache{
		Range:     rng,
		Index:     rand.Intn(len),
		ExpiredAt: time.Now().Add(m.cd),
	}
	m.randCache[coin] = tmp

	time.AfterFunc(m.cd, func() {
		if c, ok := m.randCache[coin]; ok {
			if time.Now().After(c.ExpiredAt) {
				delete(m.randCache, coin)
			}
		}
	})

	return tmp.Index
}

func readDataset(path string) (*common.MemeDataset, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ds := &common.MemeDataset{}
	err = json.Unmarshal(f, ds)
	return ds, err
}

func indexDataset(ds *common.MemeDataset) map[string]int {
	indexMap := make(map[string]int)
	for i, c := range ds.Coins {
		indexMap[c.Name] = i
	}
	return indexMap
}
