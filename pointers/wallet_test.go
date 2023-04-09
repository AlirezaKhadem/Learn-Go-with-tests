package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		output := wallet.Balance()
		expectedOutput := Bitcoin(10)

		if output != expectedOutput {
			t.Errorf(
				"got %d want %d",
				output,
				expectedOutput,
			)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		expectedOutput := Bitcoin(10)
		output := wallet.Balance()

		if output != expectedOutput {
			t.Errorf(
				"got %d want %d",
				output,
				expectedOutput,
			)
		}
	})
}
