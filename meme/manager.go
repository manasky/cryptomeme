package meme

import (
	"cryptolist/common"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Manager struct {
	ds *common.MemeDataset
	mp map[string]int
}

func New(dataset string) (*Manager, error) {
	ds, err := readDataset(dataset)
	if err != nil {
		return nil, err
	}

	return &Manager{
		ds: ds,
		mp: indexDataset(ds),
	}, nil
}

func (m *Manager) Meme(coin string, change float64) (*common.Meme, error) {
	if c, ok := m.mp[coin]; ok {
		for _, r := range m.ds.Coins[c].Ranges {
			if change <= r.Value {
				rand.Seed(time.Now().UnixNano())
				return r.Memes[rand.Intn(len(r.Memes))], nil
			}
		}
	}
	return &common.Meme{}, nil //todo
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
