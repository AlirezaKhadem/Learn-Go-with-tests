package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBitcoin := func(t testing.TB, output Bitcoin, expectedOutput Bitcoin) {
		if output != expectedOutput {
			t.Errorf(
				"got %d want %d",
				output,
				expectedOutput,
			)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		output := wallet.Balance()
		expectedOutput := Bitcoin(10)

		assertBitcoin(t, output, expectedOutput)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		expectedOutput := Bitcoin(10)
		output := wallet.Balance()

		assertBitcoin(t, output, expectedOutput)
	})
}
