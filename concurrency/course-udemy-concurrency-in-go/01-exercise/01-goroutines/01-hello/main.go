package main

import (
	"fmt"
	"time"
)

func printmsg(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	printmsg("direct call")

	// goroutine function call
	go printmsg("gorouting-1")

	// goroutine with anonymous function
	go func() {
		printmsg("goroutine-2")
	}()

	// goroutine with function value call
	fv := printmsg
	go fv("goroutine-3")

	// wait for goroutines to end
	fmt.Println("wait for goroutines")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done..")
}
