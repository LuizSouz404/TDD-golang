package main

import (
	"errors"
	"fmt"
)

func main() {}

type Bitcoin int

type Wallet struct {
	Statement Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.Statement += value
}

func (w *Wallet) Extract() Bitcoin {
	return w.Statement
}

func (w *Wallet) WithDraw(value Bitcoin) error {
	if w.Statement < value {
		return ErrorInsufficientBalance
	}
	w.Statement -= value
	return nil
}

var ErrorInsufficientBalance = errors.New("you don't have enough balance")
