package main

import "fmt"

func challenge(cnt int) (chan int, chan int){
	output := make(chan int)
	from := output

	for i := 0; i < cnt; i++ {
		input := output
		output = make(chan int)
		// make pipeline
		go func(in, out chan int) {
			for v := range in {
				out <-v
			}
			close(out)

		}(input, output)
	}

	to := output
	return from, to
}

func main()  {
	for i := 10; i < 10000000; i = i*10 {
		fmt.Printf("Challenge: %v\n", i)
		in, out := challenge(i)
		for j := 0; j < i; j++ {
			in<-j
			<-out
		}
		fmt.Printf("OK: %v\n", i)
	}
}
