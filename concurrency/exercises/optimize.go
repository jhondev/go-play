package main

import "sync"

// Optimize the sum
func sum(nums []int) int {
	t := 0
	for _, num := range nums {
		t += num
	}
	return t
}

func sumopt(nums []int) int {
	if len(nums) < 100000 {
		return sum(nums)
	}

	const split = 100
	var wg sync.WaitGroup
	wg.Add(split)

	len := len(nums) / split
	t := 0
	for i := 0; i < split; i++ {
		go func(n []int) {
			defer wg.Done()
			t += sum(n)
		}(nums[i*len : len*(i+1)])
	}
	wg.Wait()

	return t
}
