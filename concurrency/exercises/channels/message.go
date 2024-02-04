package main

import (
	"fmt"
)

func message() {
	ch := make(chan int)
	go func(a, b int) {
		t := a + b
		ch <- t
	}(1, 2)
	t := <-ch
	fmt.Println(t)
}
