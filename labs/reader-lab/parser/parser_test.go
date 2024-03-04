package parser

import (
	"bytes"
	"os"
	"reader-lab/util"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestParseReaderToType(t *testing.T) {
	data := []byte(`{"name": "John", "age": 30}`)
	person, err := ParseReader[Person](bytes.NewReader(data))
	util.FatalTestIf(t, err)
	if person.Name != "John" {
		t.Fatalf("unexpected parse: %v", person)
	}
}

func TestParseBytesToType(t *testing.T) {
	data := []byte(`{"name": "John", "age": 30}`)
	person, err := ParseBytes[Person](data)
	util.FatalTestIf(t, err)
	if person.Name != "John" {
		t.Fatalf("unexpected parse: %v", person)
	}
}

func TestDynParseBytes(t *testing.T) {
	data := []byte(`{"name": "John", "age": 30}`)
	person, err := DynParseBytes(data)
	util.FatalTestIf(t, err)
	if person["name"] != "John" {
		t.Fatalf("unexpected parse: %v", person)
	}
}

func TestParseReaderToJsonMap(t *testing.T) {
	data := []byte(`{"name": "John", "age": 30}`)
	person, err := ParseReader[JsonMap](bytes.NewReader(data))
	util.FatalTestIf(t, err)
	if (*person)["name"] != "John" {
		t.Fatalf("unexpected parse: %v", *person)
	}
}

func TestParseReaderFromFile(t *testing.T) {
	file, err := os.Open("testdata/person.json")
	util.FatalTestIf(t, err)
	defer file.Close()
	person, err := ParseReader[Person](file)
	util.FatalTestIf(t, err)
	if person.Name != "Audrey" {
		t.Fatalf("unexpected parse: %v", person)
	}
}
