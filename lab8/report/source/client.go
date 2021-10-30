package client

import (
	"fmt"
	"lab8/pkg/account"
	"math/rand"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type Accounts []account.Account

func (a Accounts) Len() int {
	return len(a)
}

func (a Accounts) Less(i, j int) bool {
	return a[i].Balance < a[j].Balance
}

func (a Accounts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Client struct {
	Name     string
	accounts Accounts
	mutex    sync.Mutex
}

func (c *Client) OpenAccount() {
	c.mutex.Lock()
	c.accounts = append(c.accounts, account.CreateAccount())
	c.mutex.Unlock()
}

func (c *Client) CloseAccount() {
	c.mutex.Lock()
	l := c.accounts.Len()
	if l > 0 {
		index := rand.Intn(l)
		c.accounts = append(c.accounts[:index], c.accounts[index+1:]...)
	}
	c.mutex.Unlock()
}

func (c *Client) SortAccounts(reverseFlag bool) {
	c.mutex.Lock()
	if reverseFlag {
		sort.Sort(sort.Reverse(c.accounts))
	} else {
		sort.Sort(c.accounts)
	}
	c.mutex.Unlock()
}

func (c *Client) PrintSumAccount() {
	c.mutex.Lock()
	sum := int64(0)
	for _, item := range c.accounts {
		sum += item.Balance
	}
	fmt.Printf("Sum of accounts: %f rub\n", account.FormatBalance(sum))
	c.mutex.Unlock()
}

func (c *Client) PrintPosNegSumAccount() {
	c.mutex.Lock()
	posSum := int64(0)
	negSum := int64(0)
	for _, item := range c.accounts {
		if item.IsNegative() {
			negSum += item.Balance
		} else {
			posSum += item.Balance
		}
	}
	fmt.Printf("Sum of positive accounts: %f rub\n", account.FormatBalance(posSum))
	fmt.Printf("Sum of negative accounts: %f rub\n", account.FormatBalance(negSum))
	c.mutex.Unlock()
}

func CreateClient() Client {
	name, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return Client{Name: name.String(), accounts: Accounts{}, mutex: sync.Mutex{}}
}
