package world

import "fmt"

// Load loads the World using the map data from the provider and
// positions aliens in random locations
// naliens: number of aliens to be positioned in the map
// randf: random function to return a pseudo-random number to position an alien
func Load(prov Provider, naliens int, maxmoves int, randf RandF) (*World, error) {
	// load map from provider
	data, err := prov.GetMap()
	if err != nil {
		return nil, err
	}
	wld := &World{}
	wld.Map = data
	wld.MaxMoves = maxmoves

	// load aliens
	wld.Aliens = make(Aliens)
	for i := 1; i <= naliens; i++ {
		randIdx := randf(len(wld.Map.Cities))
		city := wld.Map.Cities[randIdx]
		wld.Aliens[Alien(fmt.Sprint(i))] = &AlInfo{City: city, Trapped: false}

		status := wld.Map.Battlefield[city]
		status.Aliens[Alien(fmt.Sprint(i))] = false
		wld.Map.Battlefield[city] = status
	}

	return wld, nil
}
