package warehouse

import (
	"fmt"
	"lab3/pkg/product"
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
)

// Задание Информация о товарах на складе

type Warehouse struct {
	name     string
	sections []WarehouseSection
}

func (w Warehouse) Print() {
	fmt.Printf("склад %s\n", w.name)
	for s_index, s := range w.sections {
		fmt.Printf("\tсекция %d\n", s_index+1)
		s.Print("\t\t")
	}
}

type WarehouseSection struct {
	products map[product.Product]int
}

func (s WarehouseSection) Print(prefix string) {
	for product, count := range s.products {
		fmt.Printf("%s%s x%d\n", prefix, product.Name, count)
	}
}

func CreateMockWarehouse(count int, productList []product.Product) []Warehouse {
	resultNameSlice := make([]Warehouse, 0, count)

	seed := time.Now().UTC().UnixNano()
	rand.Seed(time.Now().UTC().UnixNano())
	nameGenerator := namegenerator.NewNameGenerator(seed)

	for i := 0; i <= count; i++ {
		name := nameGenerator.Generate()
		warehouseSections := CreateMockWarehouseSection(rand.Intn(5), productList)
		resultNameSlice = append(resultNameSlice, Warehouse{name, warehouseSections})
	}
	return resultNameSlice
}

func CreateMockWarehouseSection(count int, productList []product.Product) []WarehouseSection {
	resultSlice := make([]WarehouseSection, 0, count)

	for i := 1; i < count; i++ {
		section := make(map[product.Product]int)
		for _, p := range productList {
			if rand.Intn(2)-1 == 0 {
				section[p] = rand.Intn(100)
			}
		}
		resultSlice = append(resultSlice, WarehouseSection{section})
	}

	return resultSlice
}

func PrintWarehouses(warehouseList []Warehouse) {
	fmt.Println("Список складов:")
	for _, w := range warehouseList {
		fmt.Printf(" - %s\n", w.name)
	}
}

func PrettyPrintWarehouses(warehouseList []Warehouse) {
	fmt.Println("Полный список товаров по складам:")
	for _, w := range warehouseList {
		w.Print()
	}
}
