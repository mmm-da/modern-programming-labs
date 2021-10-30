package bacteria

import "fmt"

type IBacteria interface {
	tick()
	getCurrentLivetime() int
}

type BacteriaType struct {
	SplitPeriod int
	bacterias   []*Bacteria
}

func (bt BacteriaType) CountBacterias() int {
	fmt.Println("Count of bacterias", len(bt.bacterias))
	return len(bt.bacterias)
}

func NewBacteriaType(splitPeriod int) BacteriaType {
	bacterias := []*Bacteria{{currentLivetime: 1}}
	return BacteriaType{SplitPeriod: splitPeriod, bacterias: bacterias}
}

func (bt *BacteriaType) Tick() {
	newBacterias := 0
	for _, item := range bt.bacterias {
		item.tick()
		if item.getCurrentLivetime()%bt.SplitPeriod == 0 {
			newBacterias += 1
		}
	}
	for i := 0; i < newBacterias; i++ {
		bt.bacterias = append(bt.bacterias, &Bacteria{currentLivetime: 1})
	}
}

type Bacteria struct {
	currentLivetime int
}

func (b Bacteria) getCurrentLivetime() int {
	return b.currentLivetime
}

func (b *Bacteria) tick() {
	b.currentLivetime += 1
}
