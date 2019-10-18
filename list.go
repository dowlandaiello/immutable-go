// Package immutable implements a set of functional, persistent, immutable data
// collections.
package immutable

// List is an expandable, immutable data collection.
type List func(index int) interface{}

// Pair is a logical pair of interfaces.
type Pair func() (interface{}, interface{})

/* BEGIN EXPORTED METHODS */

// NewList initializes a new list from the given set of elements.
func NewList(elements ...interface{}) List {
	// Check list is empty
	if len(elements) == 0 {
		return NewList(nil) // Make an empty list
	}

	return makeList(0, nil, elements...) // Make the list
}

// Size gets the size of the list.
func (l *List) Size() int {
	return l
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// makeList generates a list from the given slice of elements.
func makeList(index int, list List, elements ...interface{}) List {
	l := func(i int) interface{} {
		if i < 0 { // Check out of bounds
			if i == -1 { // Check is identity
				return func() (interface{}, interface{}) {
					return index, elements[0] // Return the index and the element
				}
			}
			panic("index out of bounds") // Panic
		}

		if i == index { // Check is node being addressed
			return elements[0] // Return the element
		}

		return list(i) // Return the element at that index elsewhere in the list
	}

	// Check is done constructing the list
	if len(elements) == 1 {
		return l // Return the final list
	}

	return makeList(index+1, l, elements[1:]...) // Return the list after the rest of the constructing calls
}

/* END INTERNAL METHODS */
