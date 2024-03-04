package main

import (
	"fmt"
	"os"
	"reader-lab/parser"
	"reader-lab/util"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	file, err := os.Open("parser/testdata/person.json")
	util.FatalIf(err)
	defer file.Close()

	person, err := parser.ParseReader[Person](file)
	util.FatalIf(err)

	fmt.Print(*person)
}
