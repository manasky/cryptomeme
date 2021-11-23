package common

type MemeDataset struct {
	Coins []*Coin `json:"coins"`
}

type Content struct {
	X int `json:"x"`
	Y int `json:"y"`
	Text string `json:"text,omitempty"`
	Image string `json:"image,omitempty"`
}

type Meme struct {
	Image   string `json:"image"`
	Caption string `json:"caption"`
	Content []*Content `json:"content"`
}

type Range struct {
	Value float64 `json:"value"`
	Memes []*Meme `json:"memes"`
}

type Coin struct {
	Name   string   `json:"name"`
	Ranges []*Range `json:"ranges"`
}