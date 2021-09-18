package main

import "fmt"

/*
Golang - структурный язык, но он позволяет создавать иерархию структур вкладывая их друг в друга
(тем самым реализовывать аналог наследования) и прикреплять функции к структурам (создавать методы).
*/

type Hope struct {
	hopeCount int
}

//В Java статическая функция, может выполняться без контекста (инстанса класса), в golang методы нельзя вызвать без экземпляра структуры.
//Поэтому static метод execute можно заменить на функцию. 

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
	hope := Hope{}
	dream := Dream{}

	hope.taste()
	dream.taste()
	
	fmt.Println("task2")
}
