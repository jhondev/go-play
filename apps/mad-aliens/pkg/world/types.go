package world

import (
	"mad-aliens/pkg/world/directions"
	"mad-aliens/pkg/world/events"
)

// RandF is the signature of a random function that returns, as an int,
// a non-negative pseudo-random number
// in the half-open interval [0,n) from the default Source.
// It should panics if n <= 0
type RandF func(n int) int

// MAP TYPES
type City string

// Directions is a map city destinations
type Directions map[City]directions.Direction

// Well named types are used to ensure values correctness
// and better code documentation

// Alien is a well named type of string
type Alien string

// Alien is a well named type of bool
type Trapped bool

// AlInfo saves the alien info in a world
type AlInfo struct {
	City
	Trapped
}

// Aliens is a map alien info
type Aliens map[Alien]*AlInfo

// CityAliens is a map alien info in a city
type CityAliens map[Alien]Trapped

// CityStatus saves the city info in the battlefield
type CityStatus struct {
	Directions
	Aliens    CityAliens
	Destroyed bool
}

// Battlefield saves the state of the invasion
type Battlefield map[City]*CityStatus

// Map saves the information the world map including initial cities and
// the battlefield state
type Map struct {
	Cities      []City
	Battlefield // Battlefield is a map of cities status, the state of the invasion
}

// WORLD TYPES
// World
type World struct {
	Map *Map
	Aliens
	MaxMoves int
}

// PROVIDER TYPES
// Provider defines the contract to plug a new data map provider to the simulation
// this allows to have multiple data sources like apis, files and data bases
type Provider interface {
	GetMap() (*Map, error)
}

// RECORDS TYPES
// Report summarize the state of the battlefield
type Report struct {
	NAliens         int
	NCities         int
	MaxMoves        int
	TotalMoves      int
	DestroyedCities int
	DestroyedAliens int
	SurvivingAliens int
	TrappedAliens   int
	MapFormatted    []string
}

// Records defines the contract to plug a new records store to the simulation,
// this allows to have multiple data management solutions and to save
// the entire state history of a simiulation.
// The advantage of storing a state machine is to recreate an entire simulation
// or particular scenarios, giving the ability, for instance, to go back "in time"
// and having a different simulation result from a particular point of the history
type Records interface {
	LogEvent(e events.Event, wld *World) error
	FinalReport() (*Report, error)
}
