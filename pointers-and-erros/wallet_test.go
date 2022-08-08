package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Extract()
	want := 10.0

	if got != want {
		t.Errorf("Wallet\ngot: %2.f\nexpect: %2.f", got, want)
	}
}
