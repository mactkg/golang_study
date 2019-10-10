// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// cleaning
	Withdraw(Balance())
}

func TestWithdraw(t *testing.T)  {
	done := make(chan struct{})

	// init
	Deposit(500)
	fmt.Println("=", Balance())

	// Alice
	go func() {
		Withdraw(300)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Withdraw(201)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	<-done
	<-done

	if got := Balance(); got != 200 && got != 299 {
		t.Errorf("Balance = %d, want 200 or 299", got)
	}
}
