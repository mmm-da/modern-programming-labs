package main

import "fmt"

type Application interface {
	hello()
}

type ApplicationStruct struct{}

func (as ApplicationStruct) hello() {
	fmt.Println("Hello World!")
}

func callHello(a Application) {
	a.hello()
}

func main() {
	fmt.Println("task2")
	as := ApplicationStruct{}
	callHello(as)
}
