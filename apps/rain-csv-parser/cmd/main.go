package main

import (
	"log"
	"rain-csv-parser/mapper"
	"rain-csv-parser/parser"
)

func main() {
	handleErr := func(format string, err error) {
		if err != nil {
			log.Fatalf(format, err)
		}
	}

	mp, err := mapper.NewJson("./config.json")
	handleErr("error getting config: %v", err)

	err = parser.Parse("testdata/roster1.csv", mp)
	handleErr("error parsing file: %v", err)
}
