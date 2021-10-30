package main

// Автостоянка. Доступно несколько машиномест. На одном месте
// может находиться только один автомобиль. Если все места заняты, то
// автомобиль не станет ждать больше определенного времени и уедет на другую
// стоянку.

// Пусть время между появлением нового автомобиля на стоянке определяется по показательному закону распределения СВ

import (
	"lab7/pkg/parking"
	"lab7/pkg/settings"
)

func main() {
	p := make(chan int)
	stop := make(chan struct{})
	cars := parking.CreateTraffic(settings.AutoLambda, stop)
	go parking.InitPlaces(p, stop, 3)
	for {
		select {
		case car := <-cars:
			go car.Process(p, stop)
		}
	}
}
