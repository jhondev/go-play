package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		it := i // 'it' is a new variable in each iteration while 'i' is always the same varible
		wg.Add(1)
		go func() { // go func(it int) { creating a parameter is more elegant
			defer wg.Done()
			fmt.Println(it)
		}() // (i) passing argument is more elegant
	}
	wg.Wait()
}
