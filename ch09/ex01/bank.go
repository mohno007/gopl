// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type Bank struct {
	deposits chan int
	balances chan int
	withdraw chan WithdrawRequest
}

type WithdrawRequest struct {
	amount int
	result chan<- bool
}

func New() *Bank {
	b := Bank{
		deposits: make(chan int), // send amount to deposit
		balances: make(chan int), // receive balance
		withdraw: make(chan WithdrawRequest),
	}
	go b.teller() // start the monitor goroutine

	return &b
}

/// Deposit は、口座に入金します
func (b *Bank) Deposit(amount int) { b.deposits <- amount }

/// Balance は、口座の残高を返します
func (b *Bank) Balance() int { return <-b.balances }

/// Withdraw は、口座から引き出します
func (b *Bank) Withdraw(amount int) bool {
	result := make(chan bool)
	b.withdraw <- WithdrawRequest{
		amount: amount,
		result: result,
	}
	return <-result
}

func (b *Bank) teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-b.deposits:
			balance += amount
		case b.balances <- balance:
		case req := <-b.withdraw:
			if req.amount > 0 && req.amount > balance {
				req.result <- false
				continue
			}
			balance -= req.amount
			req.result <- true
		}
	}
}

//!-
