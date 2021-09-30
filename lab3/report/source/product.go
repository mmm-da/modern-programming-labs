package product

import (
	"fmt"
	"time"

	"github.com/goombaio/namegenerator"
)

type Product struct {
	Name string
}

func (p Product) Print(prefix string) {
	fmt.Printf("%s%s", p.Name, prefix)
}

func CreateMockProducts(count int) []Product {
	resultNameSlice := make([]Product, 0, count)

	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	for i := 0; i <= count; i++ {
		name := nameGenerator.Generate()
		resultNameSlice = append(resultNameSlice, Product{Name: name})
	}
	return resultNameSlice
}

func PrintProducts(productList []Product) {
	fmt.Println("Список продуктов:")
	for _, p := range productList {
		p.Print(" - ")
	}
}
