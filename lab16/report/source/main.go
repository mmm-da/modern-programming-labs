package main

import (
	"fmt"
	"lab6/pkg/bacteria"
)

// Паттерн Flyweight. Разработать систему учета процессов размножения колонии бактерий.

func main() {
	bt := bacteria.NewBacteriaType(10)
	for {
		if bt.CountBacterias() < 1000 {
			bt.Tick()
		} else {
			fmt.Println("Stop program")
			break
		}
	}
}
