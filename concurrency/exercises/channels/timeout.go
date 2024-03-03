package main

import (
	"fmt"
	"time"
)

func timeout() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "message"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
