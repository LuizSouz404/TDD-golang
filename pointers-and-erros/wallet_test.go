package main

import "testing"

func TestWallet(t *testing.T) {
	confirmStatement := func(t *testing.T, wallet Wallet, valueWant Bitcoin) {
		t.Helper()
		got := wallet.Extract()

		if got != valueWant {
			t.Errorf("Wallet\ngot: %d\nexpect: %d", got, valueWant)
		}
	}

	confirmError := func(t *testing.T, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatalf("Expected an error but none occurred")
		}

		if err != nil {
			t.Errorf("Wallet\ngot: %s\nexpect: %s", err.Error(), want.Error())
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

		err := wallet.WithDraw(Bitcoin(10))

		want := Bitcoin(10)

		confirmStatement(t, wallet, want)
		confirmError(t, err, ErrorInsufficientBalance)
	})

	t.Run("This test should not withdraw cause don't have enough balance", func(t *testing.T) {
		wallet := Wallet{Statement: Bitcoin(5)}

		err := wallet.WithDraw(Bitcoin(10))

		want := Bitcoin(10)

		confirmStatement(t, wallet, want)
		confirmError(t, err, ErrorInsufficientBalance)
	})

}
