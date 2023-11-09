package world

import (
	"strings"
)

// Clone helps to clone a world reference value to save states
func Clone(wld *World) *World {
	newMap := &Map{Cities: make([]City, 0), Battlefield: make(Battlefield)}
	newMap.Cities = append(newMap.Cities, wld.Map.Cities...)
	for k, v := range wld.Map.Battlefield {
		newMap.Battlefield[k] = v
	}

	newWld := &World{Aliens: make(Aliens), Map: newMap, MaxMoves: wld.MaxMoves}
	for k, v := range wld.Aliens {
		newWld.Aliens[k] = v
	}
	return newWld
}

// FormatMap formats a world.Map into an array of lines with the following format
// <CITY> <DIR>=<CITY> <DIR>=<CITY>
func FormatMap(wmap *Map) []string {
	lines := make([]string, 0)
	for city, status := range wmap.Battlefield {
		if status.Destroyed {
			continue
		}
		buf := strings.Builder{}
		buf.WriteString(string(city))
		for dest, dir := range status.Directions {
			buf.WriteString(" ")
			buf.WriteString(string(dir))
			buf.WriteString("=")
			buf.WriteString(string(dest))
		}
		lines = append(lines, buf.String())
	}
	return lines
}

// JoinAliens is a util function that joins a collection of aliens
// to be used in a message string
func JoinAliens(aliens CityAliens, sep string) string {
	if len(aliens) == 0 {
		return ""
	}

	buf := strings.Builder{}
	maxLen := 0
	i := 0

	for k := range aliens {
		if len(k) > maxLen {
			maxLen = len(k)
		}
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.Grow(buf.Len() + len(k) + len(sep))
		buf.WriteString(string(k))
		i++
	}

	return buf.String()
}
