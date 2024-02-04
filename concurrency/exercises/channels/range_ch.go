package main

import "fmt"

func rangech() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
		"message 5",
		"message 6"}

	// ch := make(chan string, len(messages)) // unblocking
	ch := make(chan string) // blocking
	go func() {
		defer close(ch)

		for _, msg := range messages {
			fmt.Println("sending ", msg)
			ch <- msg
		}
	}()

	for msg := range ch {
		fmt.Println("received ", msg)
	}
}
