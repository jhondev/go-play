package simulation

import (
	"mad-aliens/pkg/data/providers"
	"mad-aliens/pkg/data/records"
	"mad-aliens/pkg/world"
	"mad-aliens/pkg/world/directions"
	"math/rand"
	"testing"
)

func TestMoveAliens(t *testing.T) {
	handleErr := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	battlefield := world.Battlefield{
		"Foo": &world.CityStatus{
			Directions: world.Directions{
				"Bar": "east",
			},
			Aliens: world.CityAliens{"10": false}},
		"Bar": &world.CityStatus{
			Directions: world.Directions{
				"Foo":   "west",
				"Qu-ux": "east",
			},
			Aliens: make(world.CityAliens)},
		"Qu-ux": &world.CityStatus{
			Directions: world.Directions{
				"Bar": "east",
			},
			Aliens: world.CityAliens{"20": false}},
	}
	wmap := &world.Map{
		Cities:      []world.City{"Foo", "Bar", "Qu-ux"},
		Battlefield: battlefield,
	}
	wld := &world.World{
		Map: wmap,
		Aliens: world.Aliens{
			"10": &world.AlInfo{City: "Foo"},
			"20": &world.AlInfo{City: "Qu-ux"},
		},
		MaxMoves: 2,
	}
	rec, err := records.NewMem(wld)
	handleErr(err)
	sim := New(wld, rec, rand.Intn)
	err = sim.moveAliens()
	handleErr(err)
	if sim.world.Aliens["10"].City != sim.world.Aliens["20"].City {
		t.Fatalf("\nerror: %v is not equal to %v",
			sim.world.Aliens["10"].City,
			sim.world.Aliens["20"].City)
	}
	if len(sim.world.Map.Battlefield["Bar"].Aliens) != 2 {
		t.Fatalf("\nexpected 2 aliens in Bar city got %d", len(sim.world.Map.Battlefield["Bar"].Aliens))
	}
}

func TestSimulation_RandDest(t *testing.T) {
	handleErr := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
	tt := []struct {
		name  string
		randf world.RandF
	}{{
		name:  "no empty zero idx",
		randf: func(n int) int { return 0 },
	}, {
		name:  "no empty last idx",
		randf: func(n int) int { return n },
	}}
	dirs := world.Directions{
		"City1": directions.East,
		"City2": directions.North,
		"City3": directions.West,
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Simulation{randf: tc.randf}
			got, err := s.randDest(dirs)
			handleErr(err)

			if got == "" {
				t.Fatal("got empty")
			}
		})
	}
}

func TestMovesCounter(t *testing.T) {
	handleErr := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
	prov := providers.NewFile("./testdata/world_map_battles.txt")
	wld, err := world.Load(prov, 2, 100, rand.Intn)
	handleErr(err)
	rec, err := records.NewMem(wld)
	handleErr(err)
	sim := New(wld, rec, rand.Intn)
	handleErr(err)

	type expected struct {
		totalmoves int
	}
	tt := []struct {
		name     string
		expected expected
	}{{
		name: "move 1",
		expected: expected{
			totalmoves: 1,
		},
	}, {
		name: "move 2",
		expected: expected{
			totalmoves: 2,
		},
	}}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sim.checkBattlefield()
			if sim.checkWorld() {
				return
			}
			err = sim.moveAliens()
			handleErr(err)
			report, err := sim.records.FinalReport()
			handleErr(err)
			if report.TotalMoves != tc.expected.totalmoves {
				t.Fatalf("expected total moves: %v", tc.expected.totalmoves)
			}
		})
	}
}
