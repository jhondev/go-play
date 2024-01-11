package main

import (
	"fmt"
	"regexp"
	"sort"
)

type Supplier struct {
	name string
	cost float32
}

var Suppliers = map[string][]Supplier{
	"+31": {{name: "BeakOn", cost: 0.01}},
	"+32": {
		{name: "BeakOn", cost: 0.01},
		{name: "Swanizon", cost: 0.02},
	},
	"+33": {{name: "Swanizon", cost: 0.03}},
	"all": {{name: "TweetMobile", cost: 0.04}},
}

func main() {
	batch := []string{
		"+31000000001", // costs 1 cent1
		"+32000000001", // costs 1 cent1
		"+33000000001", // costs 3 cents3
		"+32000000002", // costs 1 cent1
		"+3@00000003",  // costs 0 cents0
		"+34000000001", // costs 4 cents4
		"+44000000001", // costs 4 cents4
		"+31000000002", // costs 1 cent1
		"+44000000002", // costs 4 cents4
		"+31000000003", // costs 1 cent1
	}

	total, cost := SenMessage("any message", batch)

	fmt.Printf("Sent %d messages at a total cost of: %v cents", total, cost)
}

func SenMessage(m string, numbers []string) (int, float32) {
	var total float32 = 0
	var sent int = 0
	for _, number := range numbers {
		pfx := number[:3]
		if !ValidatePfx(pfx) {
			continue
		}
		suppliers, ok := Suppliers[pfx]
		sent++
		if !ok {
			total += GetMinCost(Suppliers["all"])
			continue
		}
		total += GetMinCost(suppliers)
	}

	return sent, (total * 100)
}

func ValidatePfx(pfx string) bool {
	re := regexp.MustCompile(`^\+\d{2}$`)
	return re.MatchString(pfx)
}

func GetMinCost(suppliers []Supplier) float32 {
	sort.Slice(suppliers, func(i, j int) bool {
		return suppliers[i].cost < suppliers[j].cost
	})
	return suppliers[0].cost
}
