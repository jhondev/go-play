package diff

import (
	"bufio"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/blake2b"
)

// Checksum weak + strong (strong to avoid false positives)
type Checksum struct {
	Weak   uint32
	Strong string
}

// Signature consist of a list of checksums for successive fixed-size blocks.
type Signature struct {
	Checksums []Checksum
}

// WeakChecksum calculates a weak adler32 based checksum using the rollsum api
func WeakChecksum(block []byte) uint32 {
	weak := NewRollsum()
	return weak.Update(block).Sum()
}

// StrongChecksum calculates a strong blake2b checksum
func StrongChecksum(block []byte) string {
	strong := blake2b.Sum256(block)
	return hex.EncodeToString(strong[:32])
}

// Signature creates the weak + strong checksum stream signature.
// Reader contains a chunked file buffer
func (d *Diff) Signature(reader *bufio.Reader) (*Signature, error) {
	block := make([]byte, d.blockSize)

	signature := &Signature{}

	// read chunks
	for {
		bytesRead, err := reader.Read(block)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}
		weakcheck := WeakChecksum(block)
		strongcheck := StrongChecksum(block)
		checksum := Checksum{Weak: weakcheck, Strong: strongcheck}
		signature.Checksums = append(signature.Checksums, checksum)
	}

	return signature, nil
}
