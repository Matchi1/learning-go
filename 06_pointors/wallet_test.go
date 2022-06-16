package main

import (
    "testing"
)

func TestWallet(t *testing.T) {
    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))

    got := wallet.Balance()
    want := Bitcoin(10)

    if want != got {
        t.Errorf("got %s want %s", got, want)
    }
}
