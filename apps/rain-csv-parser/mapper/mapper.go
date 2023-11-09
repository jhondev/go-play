package mapper

import (
	"rain-csv-parser/mapper/providers"
	"strings"

	"golang.org/x/exp/slices"
)

type Config struct {
	Name string            `json:"name"`
	Map  map[string]string `json:"map"`
}

type Cell struct {
	Name    string
	Indexes []int
}

type Mapper struct {
	provider providers.Provider[[]Config]
	config   []Config
}

// New creates a new mapper and loads config data
func New(prov providers.Provider[[]Config]) (*Mapper, error) {
	cfg, err := prov.GetData()
	if err != nil {
		return nil, err
	}
	return &Mapper{provider: prov, config: *cfg}, nil
}

// NewJson creates a new default jsonmapper using the new builder
func NewJson(path string) (*Mapper, error) {
	return New(providers.NewJson[[]Config](path))
}

func (mp *Mapper) FindMap(header []string) (name string, cells []Cell) {
	for _, cfg := range mp.config {
		name = cfg.Name
		cells := mapcells(cfg.Map, header)
		if len(cells) > 0 {
			return name, cells
		}
	}
	return "", nil
}

func mapcells(mp map[string]string, header []string) (cells []Cell) {
	for field, fmap := range mp {
		cell := mapcell(header, field, fmap)
		if cell == nil {
			return nil
		}
		cells = append(cells, *cell)
	}
	return cells
}

func mapcell(header []string, field string, fmap string) *Cell {
	cell := Cell{Name: field}
	for _, f := range strings.Split(fmap, "+") {
		index := slices.Index(header, f)
		if index == -1 {
			return nil
		}
		cell.Indexes = append(cell.Indexes, index)
	}
	return &cell
}
