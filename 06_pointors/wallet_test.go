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
}
