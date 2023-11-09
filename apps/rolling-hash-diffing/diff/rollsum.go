package diff

// Based on the rolling checksum https://rsync.samba.org/tech_report/node3.html

// MAGIC_NUMBER is a prime number smaller than 2^16
// to handle smaller numbers using mod
const MAGIC_NUMBER = 65521

type Rollsum struct {
	count  int
	window []byte
	old    uint8
	s1, s2 uint16
}

func NewRollsum() Rollsum {
	return Rollsum{window: make([]byte, 0)}
}

// Update
func (r Rollsum) Update(data []byte) Rollsum {
	for index, char := range data {
		r.s1 += uint16(char)
		r.s2 += uint16(len(data)-index) * uint16(char)
		r.count++
	}

	r.s1 %= MAGIC_NUMBER
	r.s2 %= MAGIC_NUMBER
	return r
}

// RollIn
func (h Rollsum) RollIn(in byte) Rollsum {
	h.s1 = (h.s1 + uint16(in)) % MAGIC_NUMBER
	h.s2 = (h.s2 + h.s1) % MAGIC_NUMBER
	h.window = append(h.window, in)
	h.count++
	return h
}

// RollOut
func (r Rollsum) RollOut() Rollsum {
	if len(r.window) == 0 {
		r.count = 0
		return r
	}

	r.old = r.window[0]
	r.s1 = (r.s1 - uint16(r.old)) % MAGIC_NUMBER
	r.s2 = (r.s2 - (uint16(len(r.window)) * uint16(r.old))) % MAGIC_NUMBER
	r.window = r.window[1:]
	r.count--

	return r
}

// Sum
func (r Rollsum) Sum() uint32 {
	// base 16
	return uint32(r.s2)<<16 | uint32(r.s1)&0xFFFFF
}

func (r Rollsum) Count() int     { return r.count }
func (r Rollsum) Window() []byte { return r.window }
func (r Rollsum) Removed() uint8 { return r.old }
