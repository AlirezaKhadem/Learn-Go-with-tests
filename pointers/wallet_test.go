package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expectedOutput Bitcoin) {
		t.Helper()
		output := wallet.Balance()

		if output != expectedOutput {
			t.Errorf("got %s want %s", output, expectedOutput)
		}
	}

	assertError := func(t testing.TB, err error) {
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		expectedOutput := Bitcoin(10)
		assertBalance(t, wallet, expectedOutput)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		withdrawError := wallet.Withdraw(Bitcoin(10))
		assertError(t, withdrawError)

		expectedOutput := Bitcoin(10)
		assertBalance(t, wallet, expectedOutput)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		withdrawError := wallet.Withdraw(Bitcoin(100))

		assertError(t, withdrawError)
		assertBalance(t, wallet, startingBalance)
	})
}
