package simulation

import (
	"fmt"
	"mad-aliens/pkg/world"
	"mad-aliens/pkg/world/events"
)

type Simulation struct {
	world   *world.World
	records world.Records
	randf   world.RandF
}

func New(wld *world.World, rec world.Records, randf world.RandF) *Simulation {
	return &Simulation{
		world:   wld,
		records: rec,
		randf:   randf,
	}
}

// Run runs the simulation
func (sim *Simulation) Run() error {
	for i := 0; i < sim.world.MaxMoves; i++ {
		// check for cities status and cleanup the battlefield
		sim.checkBattlefield()
		// check for world and aliens status, break if the invasion is finished
		if sim.checkWorld() {
			break
		}
		err := sim.moveAliens()
		if err != nil {
			return err
		}
	}

	return nil
}

// checkBattlefield verifies if there are aliens fighting in the battlefield,
// cleaning up the cities and paths and updates the world map
// it returns the number of battles found
func (sim *Simulation) checkBattlefield() int {
	fmt.Printf("   👾 Checking for battlefield status:\n\n")
	battles := 0
	for cityName, status := range sim.world.Map.Battlefield {
		if len(status.Aliens) < 2 {
			continue
		}
		battles++
		fmt.Printf("      💣 A dead battle is happening in %s city\n", cityName)
		fmt.Printf("      💀 %s has been destroyed by aliens %s\n\n",
			cityName, world.JoinAliens(status.Aliens, ", "))

		sim.cleanUpBattlefield(cityName)
	}
	if battles == 0 {
		fmt.Printf("      🪧 No battles between aliens\n\n")
	}

	return battles
}

// cleanUpBattlefield updates the cities and paths in the battlefield
func (sim *Simulation) cleanUpBattlefield(city world.City) {
	sim.world.Map.Battlefield[city].Destroyed = true

	// remove paths to the city
	for _, st := range sim.world.Map.Battlefield {
		delete(st.Directions, city)
	}

	// remove city aliens from the world
	for alien := range sim.world.Map.Battlefield[city].Aliens {
		delete(sim.world.Aliens, alien)
	}

	// remove city from the battlefield
	delete(sim.world.Map.Battlefield, city)

	// save world state
	sim.records.LogEvent(events.CleanedUp, sim.world)
}

// checkWorld checks if the invasion is done.
// An invasion is done if one of the following occurs:
//   - all aliens were destroyed or there is just one alive
//   - just one alien is free
func (sim *Simulation) checkWorld() bool {
	if len(sim.world.Aliens) <= 1 {
		return true
	}
	free := 0
	for _, info := range sim.world.Aliens {
		if !info.Trapped {
			free++
		}
		if free > 1 {
			break
		}
	}

	return false
}

// moveAliens moves all aliens in the battlefield to a random destination
// based on the city paths, updates alien trapped status if necessary
func (sim *Simulation) moveAliens() error {
	fmt.Printf("      🛸 Aliens are moving\n\n")
	for alien, info := range sim.world.Aliens {
		status, ok := sim.world.Map.Battlefield[info.City]
		if !ok {
			return fmt.Errorf("error querying city %v in the battlefield", info.City)
		}
		if len(status.Directions) == 0 {
			// alien trapped
			info.Trapped = true
			status.Aliens[alien] = true
			fmt.Printf("         🪤 alien %v trapped in %v city\n\n", alien, info.City)
			continue
		}
		dest, err := sim.randDest(status.Directions)
		if err != nil {
			return err
		}
		// change alien info in the world state
		sim.world.Aliens[alien] = &world.AlInfo{City: dest, Trapped: false}
		// change cities aliens in the battlefield state
		delete(status.Aliens, alien)
		destStatus, ok := sim.world.Map.Battlefield[dest]
		if !ok {
			return fmt.Errorf("battlefield error: %s not found. Check your map source", dest)
		}
		destStatus.Aliens[alien] = false // add alien to the destination city with 'no trapped' status
	}

	// save world state
	sim.records.LogEvent(events.Moved, sim.world)
	return nil
}

// randDest uses the provided rand function to return a random destination
func (sim *Simulation) randDest(dirs world.Directions) (world.City, error) {
	if len(dirs) == 0 {
		return "", fmt.Errorf("error calculating a random dest: no directions provided")
	}
	if len(dirs) == 1 {
		for cityName := range dirs {
			return cityName, nil
		}
	}
	randIdx := sim.randf(len(dirs) - 1)
	dirIdx := 0
	for cityName := range dirs {
		if dirIdx == randIdx {
			return cityName, nil
		}
		dirIdx++
	}

	return "", fmt.Errorf("error calculating a random dest: "+
		"no result for directions: %v, randIdx: %v", dirs, randIdx)
}

func (sim *Simulation) PrintReport() {
	report, err := sim.records.FinalReport()
	if err != nil {
		fmt.Printf("error generating the final report: %v\n", err)
		return
	}
	fmt.Printf("      Simulation World\n\n")
	fmt.Printf("         aliens created: %d\n", report.NAliens)
	fmt.Printf("         cities created: %d\n", report.NCities)
	fmt.Printf("         allowed moves:  %d\n\n", report.MaxMoves)

	fmt.Printf("      Simulation Stats\n\n")
	fmt.Printf("         total moves:      %d\n", report.TotalMoves)
	fmt.Printf("         destroyed cities: %d\n", report.DestroyedCities)
	fmt.Printf("         destroyed aliens: %d\n", report.DestroyedAliens)
	fmt.Printf("         surviving aliens: %d\n", report.SurvivingAliens)
	fmt.Printf("         trapped aliens: %d\n", report.TrappedAliens)
	fmt.Printf("\n\n")

	fmt.Printf("      Aftermath Map\n\n")
	if len(report.MapFormatted) == 0 {
		fmt.Printf("         No cities left, the world was completely destroyed\n\n")
		return
	}
	for _, m := range report.MapFormatted {
		fmt.Print("         ")
		fmt.Println(m)
	}
	fmt.Printf("\n\n")
}
