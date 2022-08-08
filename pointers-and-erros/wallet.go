package main

func main() {}

type Wallet struct {
	Statement float64
}

func (w *Wallet) Deposit(value float64) {
	w.Statement += value
}

func (w *Wallet) Extract() float64 {
	return w.Statement
}
