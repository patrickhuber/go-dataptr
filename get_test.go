package dataptr_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/patrickhuber/go-dataptr"
)

func TestGet(t *testing.T) {
	type test struct {
		name     string
		obj      any
		path     string
		expected any
		err      error
	}

	tests := []test{
		{
			name:     "map",
			obj:      map[string]any{"hello": "world"},
			path:     "/hello",
			expected: "world",
		},
		{
			name:     "string_map",
			obj:      map[string]string{"hello": "world"},
			path:     "/hello",
			expected: "world",
		},
		{
			name:     "nested_map",
			obj:      map[string]any{"hello": map[string]any{"good": "world"}},
			path:     "/hello/good",
			expected: "world",
		},
		{
			name: "map_key_not_found",
			obj:  map[string]any{},
			path: "/hello",
			err:  fmt.Errorf("not found"),
		},
		{
			name:     "slice_index",
			obj:      []any{"hello", "world"},
			path:     "/0",
			expected: "hello",
		},
		{
			name: "slice_index_oob",
			obj:  []any{"hello", "world"},
			path: "/3",
			err:  fmt.Errorf("out of bounds"),
		},
		{
			name:     "constraint",
			obj:      []any{map[string]any{"hello": "world"}, map[string]any{"key": "value"}},
			path:     "/hello=world/hello",
			expected: "world",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := dataptr.Get(test.path, test.obj)
			if err != nil {
				if test.err == nil {
					t.Fatal(err)
				}
			} else if !reflect.DeepEqual(test.expected, actual) {
				t.Fatalf("actual '%s' did not match expected '%s'", actual, test.expected)
			}
		})
	}
}
