package providers

import (
	"bufio"
	"fmt"
	"mad-aliens/pkg/world"
	"mad-aliens/pkg/world/directions"
	"os"
	"strings"
)

// FileProvider will use a file as data source in the format
// <CITY> <DIR>=<CITY> <DIR>=<CITY>
// and will parse such format into the world.Map data-structure
type FileProvider struct {
	path string
}

func NewFile(path string) world.Provider {
	return &FileProvider{path: path}
}

func (fp *FileProvider) GetMap() (*world.Map, error) {
	file, err := os.Open(fp.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	wm := &world.Map{
		Cities:      make([]world.City, 0),
		Battlefield: make(world.Battlefield),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) < 2 {
			return nil, fmt.Errorf("source file bad formatted: %s", scanner.Text())
		}
		city := parts[0]
		wm.Cities = append(wm.Cities, world.City(city))
		dirs := make(world.Directions)
		for _, part := range parts[1:] {
			dirParts := strings.Split(part, "=")
			if len(dirParts) < 2 {
				return nil, fmt.Errorf("source file bad formatted %s", part)
			}
			dir := dirParts[0]
			dest := dirParts[1] // destination city in the above direction
			dirs[world.City(dest)] = directions.Direction(dir)
		}
		if _, ok := wm.Battlefield[world.City(city)]; ok {
			return nil, fmt.Errorf("duplicate city %v", city)
		}
		wm.Battlefield[world.City(city)] = &world.CityStatus{
			Directions: dirs,
			Aliens:     make(world.CityAliens),
		}
	}
	return wm, nil
}
