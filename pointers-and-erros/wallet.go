package main

func main() {}

type Bitcoin float64

type Wallet struct {
	Statement Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.Statement += value
}

func (w *Wallet) Extract() Bitcoin {
	return w.Statement
}
