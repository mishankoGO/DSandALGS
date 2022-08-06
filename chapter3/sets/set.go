package sets

// Set class
type Set struct {
	IntegerMap map[int]bool
}

// create New set method
func (set *Set) New() {
	set.IntegerMap = make(map[int]bool)
}

// AddElement method
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.IntegerMap[element] = true
	}
}

// DeleteElement method
func (set *Set) DeleteElement(element int) {
	delete(set.IntegerMap, element)
}

// ContainsElement method
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.IntegerMap[element]
	return exists
}

// InterSect method
func (set *Set) InterSect(anotherSet *Set) *Set {
	var intersection = &Set{}
	intersection.New()
	for k := range anotherSet.IntegerMap {
		if set.ContainsElement(k) {
			intersection.AddElement(k)
		}
	}
	return intersection
}

// Union method
func (set *Set) Union(anotherSet *Set) *Set {
	var union = &Set{}
	union.New()

	for k := range set.IntegerMap {
		union.AddElement(k)
	}
	for k := range anotherSet.IntegerMap {
		union.AddElement(k)
	}
	return union
}
