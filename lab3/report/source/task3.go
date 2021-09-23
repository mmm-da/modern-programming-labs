package main

import "fmt"

type AbstractCook struct{}

type Cook struct {
	ac AbstractCook
}

func (c Cook) sell() {
	fmt.Println("sell")
}

func main() {
	fmt.Println("task3")
	c := Cook{}
	c.sell()
}
