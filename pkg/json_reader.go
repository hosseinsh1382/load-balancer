package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonReader[T any] interface {
	ReadJson() (T, error)
}

type DefaultJsonReader[T any] struct {
	path string
}

func NewDefaultJsonReader[T any](path string) *DefaultJsonReader[T] {
	return &DefaultJsonReader[T]{
		path: path,
	}
}

func (j *DefaultJsonReader[T]) ReadJson() (T, error) {
	var result T
	bytes, err := os.ReadFile(j.path)
	if err != nil {
		return result, fmt.Errorf("Error reading file %s: %w", j.path, err)
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, fmt.Errorf("Error parsing file %s: %w", j.path, err)
	}
	return result, nil
}
