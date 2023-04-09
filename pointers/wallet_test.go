package pointers

import "testing"

func TestWallet(t *testing.T) {

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
}
