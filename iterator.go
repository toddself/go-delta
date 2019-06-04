package delta

// OperationType is the type of an operation
type OperationType int

const (
	// Insert operation
	Insert OperationType = iota
	// Delete operation
	Delete
	// Retain operation
	Retain
	// NOOP operation
	NOOP
)

// Iterator iterates!
type Iterator struct {
	Ops    *[]Op
	Index  int
	Offset int
}

// NewIterator creates a new iterator for use
func NewIterator(ops *[]Op) Iterator {
	i := Iterator{
		Ops:    ops,
		Index:  0,
		Offset: 0,
	}
	return i
}

// Peek returns the current operation
func (i *Iterator) Peek() Op {
	if len(*i.Ops) > i.Index {
		return (*i.Ops)[i.Index]
	}
	return Op{}
}

// PeekLength checks to see how long the current op is
func (i *Iterator) PeekLength() int {
	if len(*i.Ops) > i.Index {
		opLen, err := (*i.Ops)[i.Index].Len()
		if err != nil {
			panic("Invalid operation")
		}
		return opLen - i.Offset
	}
	return -1
}

// PeekType shows you the next operation type
func (i *Iterator) PeekType() OperationType {
	if len(*i.Ops) > i.Index {
		op := (*i.Ops)[i.Index]
		if op.Retain != 0 {
			return Retain
		}

		if op.Delete != 0 {
			return Delete
		}

		if op.Insert != nil {
			return Insert
		}
	}
	return NOOP
}

// HasNext determines if an iterator has a next operation
func (i *Iterator) HasNext() bool {
	return i.PeekLength() != -1
}
