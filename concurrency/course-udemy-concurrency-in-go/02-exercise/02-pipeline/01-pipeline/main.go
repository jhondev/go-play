package main

import "fmt"

// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	// spin goroutine
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.
	ch := square(generator(2, 3, 6, 7))

	for n := range ch {
		fmt.Println(n)
	}
}
