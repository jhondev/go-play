package diff

import (
	"fmt"

	"github.com/jhondev/rolling-hash-diffing/fs"
)

// Diff encapsulates the diff api based on the rdiff algorithm
type Diff struct {
	blockSize int
}

// New Diff
func New(blockSize int) Diff {
	return Diff{
		blockSize: blockSize,
	}
}

// Compare executes the entire expected flow from the rdiff algorithm
// to get delta diffs between 2 files.
// The ideal flow uses this steps separately in distributed scenarios
func (d Diff) Compare(name1 string, name2 string) (Delta, error) {
	// The diff algorithm consists of the following steps:
	// 1. It Splits `file 1` into a series of fixed-sized blocks of size `S` bytes
	// 2. For each of these blocks, it calculates two checksums:
	//    - a weak `rolling 32-bit checksum`.
	//    - a strong `MD4/blake2b checksum`. (to avoid false positives in case weak matches)
	// 3. It creates a checksums signature to be queried and compared.
	// 4. It searches through `file 2` to find all blocks of length `S` bytes
	//    that have the same weak and strong checksums from the signature
	//    as one of the blocks of `file 1``.
	// 5. It computes the delta

	file := fs.NewFile(d.blockSize)

	// get file1 split into blocks
	file1, err := file.Open(name1)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", name1, err)
	}

	// creates signature from chunked file
	signature, err := d.Signature(file1)
	if err != nil {
		return nil, err
	}

	// get file2 to be compared with the signature
	file2, err := file.Open(name2)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", name2, err)
	}

	// computes the delta
	delta, err := d.Delta(signature, file2)
	if err != nil {
		return nil, fmt.Errorf("error getting delta for files `%s` and `%s`: %v",
			name1, name2, err)
	}

	return delta, nil
}
