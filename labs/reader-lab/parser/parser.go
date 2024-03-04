package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type JsonMap = map[string]any

func ParseReader[T any](r io.Reader) (*T, error) {
	var result T
	var decoder = json.NewDecoder(r)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func ParseBytes[T any](b []byte) (*T, error) {
	return ParseReader[T](bytes.NewReader(b))
}

func DynParseBytes(b []byte) (JsonMap, error) {
	var result any
	err := json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	data, ok := result.(JsonMap)
	if !ok {
		return nil, fmt.Errorf("invalid json structure: %v", result)
	}
	return data, nil
}
