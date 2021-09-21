package main

import "fmt"

/*
	В golang нет необходимости реализовывать методы hashCode и equals.
	Все базовые типа языка имеют реализацию hash функции.
	Структуры состоящие из базовых типов, так же могут быть хешированы и использованы в хеш таблицах (map)

	Так же сравнение структур одного типа выполняет компилятор, структуры различного типа сравнивать нельзя.
*/

type Quest struct {
	a int32
	b int16
}

type Bus struct {
	type_var string
}

func main() {
	fmt.Println("task4")

	quest1 := Quest{a: 10, b: 20}
	quest2 := Quest{a: 10, b: 20}
	quest3 := Quest{a: 20, b: 20}

	fmt.Println(quest1 == quest2)
	fmt.Println(quest2 == quest3)

	fmt.Println("task5")

	bus1 := Bus{type_var: "bus1"}
	bus2 := Bus{type_var: "bus2"}

	fmt.Println(bus1 == bus2)
}
