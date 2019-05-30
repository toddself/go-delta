package delta

import (
	"errors"
)

// Op are the currency of the delta format
type Op struct {
	Insert     interface{}
	Retain     int
	Delete     int
	Attributes AttributeMap
}

// Len returns the length of the operation
func (o *Op) Len() (int, error) {
	if o.Retain != 0 {
		return o.Retain, nil
	}
	if o.Delete != 0 {
		return o.Delete, nil
	}
	if o.Insert != nil {
		switch v := o.Insert.(type) {
		case string:
			return len(v), nil
		case AttributeMap:
			return 1, nil
		default:
			return 0, errors.New("Inserts can only be strings or AttributeMaps")
		}
	}
	return 0, errors.New("Op contained no information")
}
