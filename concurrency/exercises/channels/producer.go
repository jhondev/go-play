package main

import (
	"fmt"
	"sync"
	"time"
)

// 1 Producer: sms
// 1 Consumer: send sms to twilio
// 1 Consumer: send sms to bird

type SMS struct {
	number  string
	message string
}

func producer() {
	messages := []SMS{
		{"+571", "message 1"},
		{"+572", "message 2"},
		{"+573", "message 3"},
		{"+574", "message 4"},
		{"+575", "message 5"},
		{"+576", "message 6"}}
	var wg sync.WaitGroup
	wg.Add(2)
	ctw := consumerTwilio(&wg)
	cbrd := consumerBird(&wg)
	produce(messages, ctw, cbrd)
	wg.Wait()
}

func consumerTwilio(wg *sync.WaitGroup) chan<- SMS {
	ch := make(chan SMS)
	go func() {
		defer wg.Done()
		for sms := range ch {
			time.Sleep(1 * time.Second)
			fmt.Println("twilio: ", sms.message)
		}
	}()
	return ch
}

func consumerBird(wg *sync.WaitGroup) chan<- SMS {
	ch := make(chan SMS)
	go func() {
		defer wg.Done()
		for sms := range ch {
			time.Sleep(2 * time.Second)
			fmt.Println("Bird: ", sms.message)
		}
	}()
	return ch
}

func produce(messages []SMS, consumers ...chan<- SMS) {
	for _, sms := range messages {
		for _, c := range consumers {
			c <- sms
		}
	}

	for _, c := range consumers {
		close(c)
	}
}
