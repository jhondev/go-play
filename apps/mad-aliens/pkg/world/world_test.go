package world_test

import (
	"mad-aliens/pkg/data/providers"
	"mad-aliens/pkg/world"
	"math/rand"
	"testing"
)

func TestWorld(t *testing.T) {
	prov := providers.NewFile("../data/providers/testdata/world_map_15.txt")
	wld, err := world.Load(prov, 10, 10, rand.Intn)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("world should save maxmoves", func(t *testing.T) {
		if wld.MaxMoves != 10 {
			t.Fatal("maxmoves is different from 10")
		}
	})

	t.Run("world aliens should be assigned to a city", func(t *testing.T) {
		for alien, info := range wld.Aliens {
			if info.City == "" {
				t.Fatalf("Alien %v is not located in any city", alien)
			}
		}
	})
}
