package delta

// AttributeMap is for storing attributes about the text
type AttributeMap map[string]interface{}

// Compose two attribute maps into a single AttributeMap
func (am AttributeMap) Compose(a, b *AttributeMap, keepNil bool) AttributeMap {
	attributes := make(AttributeMap)
	for k, v := range *b {
		if v != nil || keepNil {
			attributes[k] = v
		}
	}

	for k, v := range *a {
		_, ok := (*b)[k]
		if !ok {
			attributes[k] = v
		}
	}

	return attributes
}

// Diff two attribute maps, returning a list of items in b that don't match
// corresponding values in a (or don't exist in a)
func (am AttributeMap) Diff(a, b *AttributeMap) AttributeMap {
	var keys []string
	for k := range *a {
		keys = append(keys, k)
	}
	for k := range *b {
		for a := range keys {
			if keys[a] != k {
				keys = append(keys, k)
			}
		}
	}

	attributes := make(AttributeMap)
	for _, k := range keys {
		valA, _ := (*a)[k]
		valB, okB := (*b)[k]
		if valA != valB {
			if !okB {
				attributes[k] = nil
			} else {
				attributes[k] = valB
			}
		}
	}

	return attributes
}

// Invert a AttributeMap against a base AttributeMap
func (am AttributeMap) Invert(attr, base *AttributeMap) AttributeMap {
	baseInvert := make(AttributeMap)
	for k, baseVal := range *base {
		attrVal, ok := (*attr)[k]
		if ok && attrVal != baseVal {
			baseInvert[k] = baseVal
		}
	}
	for k, attrVal := range *attr {
		_, ok := baseInvert[k]
		baseVal, ok2 := (*base)[k]
		if !ok && ok2 && attrVal != baseVal {
			baseInvert[k] = nil
		}
	}
	return baseInvert
}

// Transform two attribute maps against each other
func (am AttributeMap) Transform(a, b *AttributeMap, priority bool) AttributeMap {
	transformed := make(AttributeMap)
	// b wins in non-priority transforms
	if !priority {
		for k, v := range *b {
			transformed[k] = v
		}
		return transformed
	}

	for k, v := range *b {
		_, ok := (*a)[k]
		if !ok {
			transformed[k] = v
		}

	}

	return transformed
}
