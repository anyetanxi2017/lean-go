package pointers

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// 在Go中，如果一个符号(例如变量、类型、函数等)是以小写符号开头，那么它在定义它的包之外就是私有的
	balance Bitcoin
}

func (c *Wallet) Balance() Bitcoin {

	return c.balance
}

// 接收者类型是 *Wallet 而不是 Wallet ,可以将其解读为 指向 wallet的指针
func (c *Wallet) Deposit(amount Bitcoin) {
	c.balance += amount
}

// var 关键字允许我们定义包的全局值
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (c *Wallet) Withdraw(amount Bitcoin) error {
	if amount > c.balance {
		return InsufficientFundsError
	}
	c.balance -= amount
	return nil
}
