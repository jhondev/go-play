package diff

import "testing"

var RollData = []byte("cloud management")

func TestUpdate(t *testing.T) {
	rollsum := NewRollsum().Update(RollData)
	check := rollsum.Sum()
	const expected = uint32(888079956)
	if check != expected {
		t.Errorf("Expected %d hash for '%s'. Got %d", expected, RollData, check)
	}
}

func TestRollIn(t *testing.T) {
	rollsum := NewRollsum()
	hash1 := rollsum.Update([]byte("cloud")).Sum()
	hash2 := rollsum.RollIn('c').RollIn('l').RollIn('o').RollIn('u').RollIn('d').Sum()

	if hash1 != hash2 {
		t.Errorf("RollIn should create same hash. Hash1: %d | Hash2: %d", hash1, hash2)
	}
}

func TestRollOut(t *testing.T) {
	rollsum := NewRollsum()
	hash1 := rollsum.Update([]byte("oud")).Sum()
	hash2 := rollsum.RollIn('c').RollIn('l').RollIn('o').RollIn('u').RollIn('d').RollOut().RollOut().Sum()

	if hash1 != hash2 {
		t.Errorf("RollOut should create same hash. Hash1: %d | Hash2: %d", hash1, hash2)
	}
}
