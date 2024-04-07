package main

import (
	"sync"
	"testing"
)

const small = 0.00000001

func TestDeposit(t *testing.T) {
	t.Run("Single Deposit", func(t *testing.T) {
		wallet := &Wallet{}
		wallet.Deposit(100.0)
		assertBalance(t, wallet, 100.0)
	})
}

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw with sufficient balance", func(t *testing.T) {
		wallet := &Wallet{balance: 20.0}
		err := wallet.Withdraw(10.0)
		assertNoError(t, err)
		assertBalance(t, wallet, 10.0)
	})
}

func TestConcurrency(t *testing.T) {
	wallet := Wallet{}
	numTransactions := 100
	amount := Bitcoin(10.0)

	var wg sync.WaitGroup
	wg.Add(numTransactions * 2)

	var mu sync.Mutex

	for i := 0; i < numTransactions; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			wallet.Deposit(amount)
			wg.Done()
		}()

		go func() {
			mu.Lock()
			defer mu.Unlock()
			err := wallet.Withdraw(amount)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	finalBalance := wallet.Balance()
	if finalBalance < -small || finalBalance > small {
		t.Errorf("final balance is not within acceptable range around 0: %f BTC", finalBalance)
	}
}

func assertBalance(t *testing.T, wallet *Wallet, expected Bitcoin) {
	t.Helper()
	if wallet.Balance() != expected {
		t.Errorf("got %f BTC, want %f BTC", wallet.Balance(), expected)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but got none")
	}
	if got.Error() != want {
		t.Errorf("got %q, want %q", got.Error(), want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("did not expect an error but got %v", got)
	}
}
