package delta

import (
	"testing"
)

func TestOp(t *testing.T) {
	op := Op{
		Retain: 10,
	}
	len, err := op.Len()
	if len != 10 {
		t.Errorf("Length of op should be 10, got %v", len)
	}
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	op = Op{
		Delete: 11,
	}
	len, err = op.Len()
	if len != 11 {
		t.Errorf("Length of op should be 11, got %v", len)
	}
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	op = Op{
		Insert: "Hello",
	}
	len, err = op.Len()
	if len != 5 {
		t.Errorf("Length of insert should be 5 (hello), got %v", len)
	}
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	op = Op{
		Insert: AttributeMap{"bold": true},
	}
	len, err = op.Len()
	if len != 1 {
		t.Errorf("Attribute operation is length 1, got %v", len)
	}
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	op = Op{
		Retain: 0,
	}
	len, err = op.Len()
	if err == nil {
		t.Errorf("Should be error, got %v", len)
	}

	op = Op{
		Insert: 5,
	}
	len, err = op.Len()
	if err == nil {
		t.Errorf("Should be error, got %v", len)
	}
}
