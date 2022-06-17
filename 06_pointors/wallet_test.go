package main

import (
    "testing"
)

func TestWallet(t *testing.T) {
    t.Run("depositing", func (t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        assertBalance(t, wallet, Bitcoin(10))
    })
    t.Run("withdraw with funds", func (t *testing.T) {
        wallet := Wallet{balance: Bitcoin(10)}
        err := wallet.Withdraw(Bitcoin(10))

        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(0))
    })
    t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}

func assertBalance(t testing.TB, wallet Wallet, bitcoin Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    want := bitcoin
    if want != got {
        t.Errorf("got %s want %s", got, bitcoin)
    }
}

func assertNoError (t testing.TB, got error) {
    t.Helper()
    if got != nil {
        t.Fatal("get an error but didn't want one")
    }
}

func assertError (t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
