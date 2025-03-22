package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
