package main

import "fmt"

var a = [...]int{1, 2, 3, 4, 5}
var s1 = make([]int, len(a), (len(a)+1)*2)
var s2 = []int{}

func grow() {
	println("init")
	Print()

	println("s1 = a[:]")
	s1 = a[:]
	Print()

	println("s2 = append")
	s2 = append(s1, 6)
	Print()

	println("s1[0] modified")
	s1[0] = 50
	Print()
}

func Print() {
	fmt.Print("Array  : ", a)
	fmt.Println(" | cap: ", cap(a))

	fmt.Print("Slice 1: ", s1)
	fmt.Println(" | cap: ", cap(s1))

	fmt.Print("Slice 2: ", s2)
	fmt.Println(" | cap: ", cap(s2))

	fmt.Println()
}
