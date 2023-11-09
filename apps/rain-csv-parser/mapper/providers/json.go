package providers

import (
	"encoding/json"
	"os"
)

type JsonProvider[T any] struct {
	path string
}

func NewJson[T any](path string) Provider[T] {
	return JsonProvider[T]{path: path}
}

func (prov JsonProvider[T]) GetData() (*T, error) {
	content, err := os.ReadFile(prov.path)
	if err != nil {
		return nil, err
	}
	var data T
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
