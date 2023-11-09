package records

import (
	"mad-aliens/pkg/world"
	"mad-aliens/pkg/world/events"
)

type Record struct {
	event events.Event
	world *world.World
}

// MemRecords struct aims to store all the events and state of an invasion
// in memory, not recommended for big simulations, you can use others
// implementations of the records store or create your own
type MemRecords struct {
	worldOG  *world.World
	current  *world.World
	records  []Record
	totalmvs int
}

func NewMem(wld *world.World) (world.Records, error) {
	rec := &MemRecords{}
	rec.worldOG = world.Clone(wld)
	rec.current = rec.worldOG
	rec.records = []Record{{event: events.Loaded, world: rec.worldOG}}
	return rec, nil
}

func (rec *MemRecords) LogEvent(e events.Event, wld *world.World) error {
	rec.current = world.Clone(wld)
	rec.records = append(rec.records, Record{event: e, world: rec.current})
	// track moves for the final report
	if e == events.Moved {
		rec.totalmvs++
	}
	return nil
}

func (rec *MemRecords) FinalReport() (*world.Report, error) {
	r := &world.Report{
		NAliens:         len(rec.worldOG.Aliens),
		NCities:         len(rec.worldOG.Map.Cities),
		SurvivingAliens: len(rec.current.Aliens),
		MaxMoves:        rec.worldOG.MaxMoves,
		TotalMoves:      rec.totalmvs,
	}
	r.DestroyedAliens = r.NAliens - r.SurvivingAliens
	r.DestroyedCities = r.NCities - len(rec.current.Map.Battlefield)
	r.MapFormatted = world.FormatMap(rec.current.Map)
	for _, info := range rec.current.Aliens {
		if info.Trapped {
			r.TrappedAliens++
		}
	}
	return r, nil
}
