package main

import (
	"fmt"

	"github.com/jhondev/rolling-hash-diffing/diff"
)

func main() {
	file1Path := "testfiles/scores1.txt"
	file2Path := "testfiles/scores2.txt"
	blockSize := 16

	diff := diff.New(blockSize)
	delta, err := diff.Compare(file1Path, file2Path)
	if err != nil {
		panic(fmt.Sprintf("error comparing files: %v", err))
	}

	// Print differences
	for i, diff := range delta {
		fmt.Printf("Difference in index %d: '%s'", i, string(diff.Data))
	}

	// The diff api let us use the different steps used in the compare function
	// separately. That way we can take advantage of the real use cases scenarios of
	// the rsync algorithm, for example a distributed scenario.
	// You can generate the signature of one file in one machine and
	// use it in another machine to compare to a second file, compute the delta and update
	// the second file
	// https://rsync.samba.org/tech_report/node2.html
	// https://librsync.github.io/page_rdiff.html
}
