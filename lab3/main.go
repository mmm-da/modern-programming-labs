package main

// Задание Информация о товарах на складе

import (
	"lab3/pkg/product"
	"lab3/pkg/warehouse"
)

func main() {
	productList := product.CreateMockProducts(10)
	warehouseList := warehouse.CreateMockWarehouse(5, productList)
	product.PrintProducts(productList)
	warehouse.PrintWarehouses(warehouseList)
	warehouse.PrettyPrintWarehouses(warehouseList)
}
