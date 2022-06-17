package main

import (
    "testing"
)

func TestWallet(t *testing.T) {
    assertBalance := func (t testing.TB, wallet Wallet, bitcoin Bitcoin) {
        got := wallet.Balance()
        want := bitcoin
        if want != got {
            t.Errorf("got %s want %s", got, bitcoin)
        }
    }
    assertError := func(t testing.TB, got error, want string) {
        t.Helper()
        if got == nil {
            t.Fatal("didn't get an error but wanted one")
        }

        if got.Error() != want {
            t.Errorf("got %q, want %q", got, want)
        }
    }
    t.Run("depositing", func (t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })
    t.Run("withdrawing", func (t *testing.T) {
        wallet := Wallet{balance: Bitcoin(10)}
        wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(0))
    })
    t.Run("withdraw insufficient funds", func(t *testing.T) {
	startingBalance := Bitcoin(20)
	wallet := Wallet{startingBalance}
	err := wallet.Withdraw(Bitcoin(100))

    assertError(t, err, "cannot withdraw, insufficient funds")
	assertBalance(t, wallet, startingBalance)
})
}
