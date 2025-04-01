package dataptr_test

import (
	"reflect"
	"testing"

	"github.com/patrickhuber/go-dataptr"
)

func TestParse(t *testing.T) {
	type test struct {
		name string
		str  string
		ptr  dataptr.DataPointer
	}
	tests := []test{

		{"name", "name", dataptr.DataPointer{
			Segments: []dataptr.Segment{
				dataptr.Element{
					Name: "name",
				},
			},
		}},
		{"index", "0", dataptr.DataPointer{
			Segments: []dataptr.Segment{
				dataptr.Index{
					Index: 0,
				},
			},
		}},
		{"constraint", "key=value", dataptr.DataPointer{
			Segments: []dataptr.Segment{
				dataptr.Constraint{
					Key:   "key",
					Value: "value",
				},
			},
		}},
		{"multi name", "parent/child", dataptr.DataPointer{
			Segments: []dataptr.Segment{
				dataptr.Element{
					Name: "parent",
				},
				dataptr.Element{
					Name: "child",
				},
			},
		}},
		{"name constraint", "name/key=value", dataptr.DataPointer{
			Segments: []dataptr.Segment{
				dataptr.Element{
					Name: "name",
				},
				dataptr.Constraint{
					Key:   "key",
					Value: "value",
				},
			},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := dataptr.Parse(test.str)
			if err != nil {
				t.Fatalf("excected err to be nil")
			}
			if !reflect.DeepEqual(test.ptr, actual) {
				t.Fatalf("expcected to equal actual")
			}
		})
	}
}
