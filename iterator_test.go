package delta

import "testing"

func TestIterator(t *testing.T) {
	ops := []Op{Op{Insert: "Hello world"}}
	itr := NewIterator(&ops)
	if itr.Index != 0 {
		t.Errorf("Newly initialized iterators index should be 0, got %v", itr.Index)
	}
	if itr.Offset != 0 {
		t.Errorf("Offset should be 0, got %v", itr.Offset)
	}
	op := (*itr.Ops)[0]
	if op.Insert != "Hello world" {
		t.Errorf("Op should be hello world, got %v", op.Insert)
	}

	op = itr.Peek()
	if op.Insert != "Hello world" {
		t.Errorf("peek() should return hello world, got %v", op.Insert)
	}

	ty := itr.PeekType()
	if ty != Insert {
		t.Errorf("peektype() should return insert, got %v", ty)
	}

	len := itr.PeekLength()
	if len != 11 { // length of "Hello world"
		t.Errorf("peeklength() should be 11, got %v", len)
	}

	next := itr.HasNext()
	if !next {
		t.Errorf("HasNext should be true, got %v", next)
	}
}
