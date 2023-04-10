package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		expectedOutput := Bitcoin(10)
		assertBalance(t, wallet, expectedOutput)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		withdrawError := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, withdrawError)

		expectedOutput := Bitcoin(10)
		assertBalance(t, wallet, expectedOutput)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		withdrawError := wallet.Withdraw(Bitcoin(100))

		assertError(t, withdrawError, InsufficientFundsError)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expectedOutput Bitcoin) {
	t.Helper()
	output := wallet.Balance()

	if output != expectedOutput {
		t.Errorf("got %s want %s", output, expectedOutput)
	}
}

func assertError(t testing.TB, err error, expectedError error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}

	if err != expectedError {
		t.Errorf(
			"got %q, want %q",
			err.Error(),
			expectedError.Error(),
		)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
