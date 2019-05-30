package iterator

import "github.com/toddself/go-delta"

// Iterator iterates
type Iterator struct {
	Ops    []delta.Op
	Index  int
	Offset int
}

// New creates a new iterator for use
func New() Iterator {
	i := Iterator{
		Ops:    make([]Op),
		Index:  0,
		Offset: 0,
	}
	return i
}
