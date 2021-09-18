package main

import "fmt"

type Hope struct {
	hopeCount int
}

func execute() {
	fmt.Println("Hope executed")
}

func (h Hope) taste() {
	execute()
	fmt.Println("Hope tasted")
}

type Dream struct {
	dream float32
}

func (d Dream) taste() float32 {
	execute()
	return d.dream
}

func main() {
	Hope
}
