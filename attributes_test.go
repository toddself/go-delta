package delta

import "testing"

func TestCompose (t *testing.T) {
	a := AttributeMap{
		"bold": true,
		"color": "#FFF",
	}
	b := AttributeMap{
		"italic": true,
	}
	composed := a.Compose(&a, &b, false)
	bold, ok := composed["bold"]
	if !ok || bold != true {
		t.Errorf("Bold should be true, got %v", bold)
	}
}