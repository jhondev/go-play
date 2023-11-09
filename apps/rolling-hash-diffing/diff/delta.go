package diff

import (
	"bufio"
	"io"
)

// Block is the file block info
type Block struct {
	Start int
	End   int
	Data  []byte
}

// Delta differences block stream
type Delta map[int]Block

// Indexes is the hash map to perform O(1) lookups
type Indexes map[uint32]map[string]int

// Delta gets the delta diff between 2 files based on the rsync and rdiff algorithms
// using a signature from the first file and the second file buffer
// deltas consist of a differences block stream.
func (d *Diff) Delta(signature *Signature, reader *bufio.Reader) (Delta, error) {
	weakSum := NewRollsum()
	delta := make(Delta)
	indexes := WeakIndexes(signature)

	var diffs []byte

	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// add byte to the data structure
		weakSum = weakSum.RollIn(b)
		// continue adding until reach window size
		if weakSum.Count() < d.blockSize {
			continue
		}

		// slice window
		if weakSum.Count() > d.blockSize {
			weakSum = weakSum.RollOut()
			removed := weakSum.Removed()
			diffs = append(diffs, removed)
		}

		// rolling hash lookup
		index := Lookup(weakSum.Window(), indexes, weakSum.Sum())
		if ^index != 0 {
			diffBlock := d.block(index, diffs)
			delta[index] = diffBlock
		}
	}

	return delta, nil
}

// WeakIndexes creates an indexes map for weak and strong checksums
func WeakIndexes(signature *Signature) Indexes {
	indexes := make(Indexes)
	for i, check := range signature.Checksums {
		indexes[check.Weak] = map[string]int{check.Strong: i}
	}

	return indexes
}

// Lookup for matches using weak and strong checksums from indexes map
func Lookup(window []byte, indexes Indexes, weak uint32) int {
	// If weak checksum exists, calculate strong checksum to avoid false positives
	if strongs, found := indexes[weak]; found {
		strong := StrongChecksum(window)
		if _, ok := strongs[strong]; ok {
			return strongs[strong]
		}
	}
	return -1
}

// block creates a block with data in a range position
func (d *Diff) block(index int, data []byte) Block {
	return Block{
		Start: (index * d.blockSize),
		End:   ((index * d.blockSize) + d.blockSize),
		Data:  data,
	}
}
