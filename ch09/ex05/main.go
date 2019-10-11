package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	done := make(chan struct{})

	f := func(i, o chan int, done <-chan struct{}) {
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case v:= <-i:
				o <- v+1
			}
		}
		v := <-i
		o <- v+1
	}

	go f(in, out, done)
	go f(out, in, done)
	in<-0

	<-time.After(time.Second)
	fmt.Println(<-out)
}
