package delta

import (
	"testing"
)

func TestCompose(t *testing.T) {
	a := AttributeMap{
		"bold":  true,
		"color": "#FFF",
	}
	b := AttributeMap{
		"italic": true,
	}
	composed := a.Compose(&a, &b, false)
	if composed["bold"] != true {
		t.Errorf("Bold should be true, got %v", composed["bold"])
	}
	if composed["italic"] != true {
		t.Errorf("Italic should be true, got %v", composed["italic"])
	}
	if composed["color"] != "#FFF" {
		t.Errorf("Color should be #FFF, got %v", composed["color"])
	}

	b = AttributeMap{
		"bold":  false,
		"color": "#AAA",
	}
	composed = a.Compose(&a, &b, false)
	if composed["bold"] != false {
		t.Errorf("Bold should be false, got %v", composed["bold"])
	}
	if composed["color"] != "#AAA" {
		t.Errorf("Color should be #AAA, got %v", composed["color"])
	}

	b = AttributeMap{
		"bold": nil,
	}
	composed = a.Compose(&a, &b, false)
	_, ok := composed["bold"]
	if ok {
		t.Errorf("Bold should be removed")
	}
	if composed["color"] != "#FFF" {
		t.Errorf("Color should be #FFF, got %v", composed["color"])
	}

	b = AttributeMap{
		"bold":  nil,
		"color": nil,
	}
	composed = a.Compose(&a, &b, false)
	if len(composed) > 0 {
		t.Errorf("All items should have been removed, but saw %v", composed)
	}

	b = AttributeMap{
		"italics": nil,
	}
	composed = a.Compose(&a, &b, false)
	_, ok = composed["italics"]
	if ok {
		t.Errorf("Italics should have not been composed since it was nil")
	}

	b = AttributeMap{
		"italics": nil,
	}
	composed = a.Compose(&a, &b, true)
	italics, ok := composed["italics"]
	if !ok || italics != nil {
		t.Errorf("Italics should have been composed since we said keep nil")
	}
}

func TestDiff(t *testing.T) {
	a := AttributeMap{
		"bold":  true,
		"color": "red",
	}

	diff := a.Diff(&a, &a)
	if len(diff) > 0 {
		t.Errorf("diff AttributeMap should be empty when comparing the same object")
	}

	b := AttributeMap{
		"bold":   true,
		"italic": true,
		"color":  "red",
	}

	diff = a.Diff(&a, &b)
	if diff["italic"] != true {
		t.Errorf("italic should be true, got %v", diff["italic"])
	}
	if len(diff) > 1 {
		t.Errorf("diff should have one attribute, has %v", len(diff))
	}

	b = AttributeMap{
		"bold": true,
	}

	diff = a.Diff(&a, &b)
	if diff["color"] != nil {
		t.Errorf("color should be red, got %v, (%v)", diff["color"], diff)
	}
	if len(diff) > 1 {
		t.Errorf("diff should have one attribute, has %v (%v)", len(diff), diff)
	}

	b = AttributeMap{
		"bold":  true,
		"color": "blue",
	}
	diff = a.Diff(&a, &b)
	if diff["color"] != "blue" {
		t.Errorf("color should be blue, got %v", diff["color"])
	}
}

func TestInvert(t *testing.T) {
	attr := AttributeMap{
		"bold": true,
	}
	base := AttributeMap{
		"italic": true,
	}
	invert := attr.Invert(&attr, &base)
	if invert["bold"] != nil {
		t.Errorf("Expected bold to be nil, got %v", invert["bold"])
	}
	if len(invert) > 1 {
		t.Errorf("Expected only a single property, got %v", len(invert))
	}

	attr = AttributeMap{
		"bold": nil,
	}
	base = AttributeMap{
		"bold": true,
	}
	invert = attr.Invert(&attr, &base)
	if invert["bold"] != true {
		t.Errorf("expected bold to be true, got %v", invert["bold"])
	}

	attr = AttributeMap{
		"color": "red",
	}
	base = AttributeMap{
		"color": "blue",
	}
	invert = attr.Invert(&attr, &base)
	if invert["color"] != "blue" {
		t.Errorf("expected color to be blue, got %v", invert["color"])
	}

	attr = AttributeMap{
		"color": "red",
	}
	base = AttributeMap{
		"color": "red",
	}
	invert = attr.Invert(&attr, &base)
	if len(invert) > 0 {
		t.Errorf("Invert map should be empty but has %v", len(invert))
	}
	attr = AttributeMap{
		"bold":   true,
		"italic": nil,
		"color":  "red",
		"size":   "12px",
	}
	base = AttributeMap{
		"font":   "serif",
		"italic": true,
		"color":  "blue",
		"size":   "12px",
	}
	invert = attr.Invert(&attr, &base)
	if invert["bold"] != nil {
		t.Errorf("bold should be nil, got %v", invert["bold"])
	}
	if invert["italic"] != true {
		t.Errorf("italic should be true, got %v", invert["italic"])
	}
	if invert["color"] != "blue" {
		t.Errorf("color should be blue, got %v", invert["color"])
	}
	if len(invert) > 3 {
		t.Errorf("invert should have 3 items, got %v", len(invert))
	}
}

func TestTransform(t *testing.T) {
	a := AttributeMap{
		"bold":  true,
		"color": "red",
		"font":  nil,
	}

	b := AttributeMap{
		"color":  "blue",
		"font":   "serif",
		"italic": true,
	}

	transform := a.Transform(&a, &b, true)
	if transform["italic"] != true {
		t.Errorf("italic should have been true, got %v", transform["italic"])
	}
	if len(transform) > 1 {
		t.Errorf("transform should have one property, got %v", len(transform))
	}

	transform = a.Transform(&a, &b, false)
	if transform["color"] != "blue" {
		t.Errorf("color should have been blue, got %v", transform["color"])
	}
	if transform["font"] != "serif" {
		t.Errorf("font should have been serif , got %v", transform["font"])
	}
	if transform["italic"] != true {
		t.Errorf("italic should have been true, got %v", transform["italic"])
	}
}
