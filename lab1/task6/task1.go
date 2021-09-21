package main

import "fmt"

/*
	Вариант f
	interface Корабль <- abstract class Военный Корабль <- class Авианосец.
*/

type Ship interface {
	float()
}

type Battleship struct {
	mainGunGauge int
}

func (b Battleship) float() {
	fmt.Println("Battleship is floating")
}

// В golang нет наследования, структуры вкладываются друг в друга.

type Carrier struct {
	Battleship
	airstripCount int
}

func (c Carrier) float() {
	fmt.Println("Carier is floating")
}

func checkShip(s Ship) {
	fmt.Println("Ship ok!")
}

func main() {
	c := Carrier{airstripCount: 2}

	// Полиморфизм как в Java, C# и C++ тоже отсутствует.
	c.Battleship.float()
	c.float()

	checkShip(c)
}
