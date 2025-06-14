package dataptr

import (
	"fmt"
)

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

func Get(path string, obj any) (any, error) {
	dataPtr, err := Parse(path)
	if err != nil {
		return nil, err
	}
	return get(dataPtr, obj)
}

func get(dataPtr DataPointer, obj any) (any, error) {
	var current = obj
	for _, seg := range dataPtr.Segments {
		switch s := seg.(type) {

		case Index:
			// current must be a slice
			slice, ok := current.([]any)
			if !ok {
				return nil, fmt.Errorf("index segments require a slice object. Found %T", current)
			}
			if s.Index >= len(slice) {
				return nil, fmt.Errorf("index %d is greater than the slice length %d", s.Index, len(slice))
			}
			if s.Index < 0 {
				return nil, fmt.Errorf("index %d is less than zero", s.Index)
			}
			current = slice[s.Index]

		case Constraint:
			// check the key to be sure
			k, ok := s.Key.(string)
			if !ok {
				return nil, fmt.Errorf("constraint segments require a key of type 'string' found '%T", s.Key)
			}

			// current must be a slice
			slice, ok := current.([]any)
			if !ok {
				return nil, fmt.Errorf("constraint segments require a slice object. Found %T", current)
			}

			// find the matching element
			found := false
			for _, elem := range slice {
				// if elem is not a map[string]any return false
				m, ok := elem.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("constraint segments require each element to be a map[string]any. Found %T", elem)
				}

				// if we found a match, set current and break the search
				_, ok = m[k]
				if ok {
					found = true
					current = elem
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("unable to find element with key %v", s.Key)
			}

		case Key:
			// current must be a map
			m, ok := current.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("key segments require a map object. Found %T", current)
			}
			// key must be a string (for now)
			k, ok := s.Key.(string)
			if !ok {
				return nil, fmt.Errorf("key segments require a string key object. Found %T", s.Key)
			}
			v, ok := m[k]
			if !ok {
				return nil, fmt.Errorf("unable to find key '%s' in object", k)
			}
			current = v
		}
	}
	return current, nil
}

func GetAs[T any](path string, obj any) (T, error) {
	result, err := Get(path, obj)
	var zero T
	if err != nil {
		return zero, err
	}
	t, ok := result.(T)
	if !ok {
		return zero, fmt.Errorf("unable to cast %T to %T", result, zero)
	}
	return t, nil
}
