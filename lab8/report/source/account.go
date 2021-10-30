package account

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	Name     string
	Balance  int64
	isLocked bool
}

func (a Account) IsNegative() bool {
	return a.Balance > 0
}

func FormatBalance(balance int64) float32 {
	return float32(balance) / 100
}

func (a *Account) SetLock(flag bool) {
	a.isLocked = flag
}

func CreateAccount() Account {
	name, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	newAccount := Account{Name: name.String(), Balance: rand.Int63(), isLocked: false}
	return newAccount
}
