package main

import "testing"

func TestWallet(t *testing.T) {
	confirmStatement := func(t *testing.T, wallet Wallet, valueWant Bitcoin) {
		got := wallet.Extract()

		if got != valueWant {
			t.Errorf("Wallet\ngot: %d\nexpect: %d", got, valueWant)
		}
	}

	t.Run("This test should deposit 10 BTC in test wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		confirmStatement(t, wallet, want)
	})

	t.Run("This test should withdraw 10 BTC in test wallet", func(t *testing.T) {
		wallet := Wallet{Statement: Bitcoin(20)}

		wallet.WithDraw(Bitcoin(10))

		want := Bitcoin(10)

		confirmStatement(t, wallet, want)
	})

}
