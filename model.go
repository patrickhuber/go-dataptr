package dataptr

import (
	"fmt"
	"strconv"
	"strings"
)

type DataPointer struct {
	Segments []Segment
}

type Segment interface {
	segment()
	fmt.Stringer
}

type Constraint struct {
	Segment
	Value any
	Key   any
}

type Key struct {
	Segment
	Key any
}

type Index struct {
	Segment
	Index int
}

func (dp DataPointer) String() string {
	builder := strings.Builder{}
	for i, seg := range dp.Segments {
		if i > 0 {
			builder.WriteRune('/')
		}
		builder.WriteString(seg.String())
	}
	return builder.String()
}

func (c Constraint) String() string {
	return fmt.Sprintf("%s=%s", c.Key, c.Value)
}

func (e Key) String() string {
	return fmt.Sprintf("%v", e.Key)
}

func (i Index) String() string {
	return strconv.Itoa(i.Index)
}
