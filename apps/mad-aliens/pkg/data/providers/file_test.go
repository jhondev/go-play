package providers_test

import (
	"mad-aliens/pkg/data/providers"
	"testing"
)

func TestFileData(t *testing.T) {
	prov := providers.NewFile("./testdata/world_map_15.txt")
	wmap, err := prov.GetMap()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Run("map should have cities", func(t *testing.T) {
		if len(wmap.Cities) == 0 {
			t.Fatal("no cities")
		}
	})

	t.Run("map cities should have directions", func(t *testing.T) {
		for city, st := range wmap.Battlefield {
			if len(st.Directions) == 0 {
				t.Fatalf("city %v has no directions", city)
			}
		}
	})
}

func TestDuplicatesError(t *testing.T) {
	prov := providers.NewFile("./testdata/world_map_duplicates.txt")
	_, err := prov.GetMap()
	if err == nil {
		t.Fatal("get map should return error for duplicates")
	}
}
