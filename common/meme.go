package common

type MemeDataset struct {
	Coins []*Coin `json:"coins"`
}

type Meme struct {
	Image   string `json:"image"`
	Caption string `json:"caption"`
}

type Range struct {
	Value float64 `json:"value"`
	Memes []*Meme `json:"memes"`
}

type Coin struct {
	Name   string   `json:"name"`
	Ranges []*Range `json:"ranges"`
}