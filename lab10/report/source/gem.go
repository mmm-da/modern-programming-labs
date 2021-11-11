package gem

import "github.com/brianvoe/gofakeit/v6"

type Gem struct {
	Name         string    `json:"name"`
	Preciousness bool      `json:"preciousness"`
	Origin       string    `json:"origin"`
	Visual       GemVisual `json:"visual"`
	Value        int       `json:"value"`
}

type GemVisual struct {
	Color       string  `json:"color"`
	Opacity     float32 `json:"opacity"`
	CountOfEdge int     `json:"edges"`
}

func CreateMockGem() Gem {
	name := gofakeit.UUID()
	preciousness := gofakeit.Bool()
	origin := gofakeit.Country()
	value := gofakeit.Number(1, 1000)

	color := gofakeit.Color()
	opacity := gofakeit.Number(0, 100) / 100
	countOfEdge := gofakeit.Number(3, 50)

	visual := GemVisual{Color: color, Opacity: float32(opacity), CountOfEdge: countOfEdge}

	result := Gem{Name: name, Preciousness: preciousness, Origin: origin, Visual: visual, Value: value}
	return result
}
