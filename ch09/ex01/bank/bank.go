// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var success = make(chan bool)
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Withdraw(amount int) bool {
	if amount < 0 {
		return false
	}

	deposits <- -amount
	return <-success
}
func Deposit(amount int) bool {
	deposits <- amount
	return <-success
}
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			if balance+amount >= 0 {
				balance += amount
				success <- true
			} else {
				success <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
