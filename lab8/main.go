package main

import (
	"lab8/pkg/client"
	"math/rand"
)

// Счета. Клиент может иметь несколько счетов в банке. Учитывать
// возможность блокировки/разблокировки счета. Реализовать поиск и
// сортировку счетов. Вычисление общей суммы по счетам. Вычисление суммы
// по всем счетам, имеющим положительный и отрицательный балансы
// отдельно.

func main() {
	c := client.CreateClient()
	for {
		action := rand.Intn(7)
		switch action {
		case 1:
			go c.OpenAccount()
		case 2:
			go c.SortAccounts(false)
		case 3:
			go c.SortAccounts(true)
		case 4:
			go c.PrintPosNegSumAccount()
		case 5:
			go c.PrintSumAccount()
		case 6:
			go c.CloseAccount()
		}
	}
}
