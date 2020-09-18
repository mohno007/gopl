// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/mohno007/gopl/ch09/ex01"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	account := bank.New()

	// Alice
	go func() {
		account.Deposit(200)
		fmt.Println("=", account.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		account.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := account.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBankWithdraw(t *testing.T) {
	done := make(chan struct{})

	account := bank.New()

	// Alice
	go func() {
		account.Deposit(200)
		fmt.Println("=", account.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		account.Withdraw(200)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got := account.Balance(); got != 0 && got != 200 {
		t.Errorf("Balance = %d, want 0 or 200", got)
	}
}

func TestBankWithdrawFail(t *testing.T) {
	account := bank.New()

	// Alice
	account.Deposit(200)
	fmt.Println("=", account.Balance())

	// Bob
	if got, want := account.Withdraw(201), false; got != want {
		t.Errorf("Balance = %v, want %v", got, want)
	}

	if got, want := account.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBankWithdrawOk(t *testing.T) {
	account := bank.New()

	// Alice
	account.Deposit(200)
	fmt.Println("=", account.Balance())

	// Bob
	if got, want := account.Withdraw(200), true; got != want {
		t.Errorf("Balance = %v, want %v", got, want)
	}

	if got, want := account.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
