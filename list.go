// Package immutable implements a set of functional, persistent, immutable data
// collections.
package immutable

// List is an expandable, immutable data collection.
type List func(index int, callback ...func(index int, element interface{}) interface{}) interface{}

// pair is a logical pair of interfaces.
type pair func() (interface{}, interface{})

/* BEGIN EXPORTED METHODS */

// NewList initializes a new list from the given set of elements.
func NewList(elements ...interface{}) List {
	// Check list is empty
	if len(elements) == 0 {
		return NewList(nil) // Make an empty list
	}

	return makeList(0, nil, elements...) // Make the list
}

// ForEach performs the given callback for each of the items in the list.
func (l List) ForEach(callback func(index int, value interface{}) interface{}) {
	(l)(0, callback) // Do the callback for each of the items in the list
}

// Set sets the value of an element in a list at the given index.
func (l List) Set(index int, value interface{}) List {
	return l(-3, func(i int, element interface{}) interface{} {
		// Check at correct node
		if i == index {
			return value // Return the proper value
		}

		return nil // Nothing to return
	}).(List) // Set the element
}

// Push pushes an element to the list.
func (l List) Push(element interface{}) List {
	return makeList(l.Size(), l, element) // Push the element to the list
}

// Pop removes an element from the list.
func (l List) Pop() List {
	return l(-2).(List) // Return the popped list
}

// Size gets the size of the list.
func (l List) Size() int {
	size, _ := l(-1).(pair)() // Get the size of the list

	return size.(int) + 1 // Return the size of the list
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// makeList generates a list from the given slice of elements.
func makeList(index int, list List, elements ...interface{}) List {
	l := func(i int, callback ...func(index int, element interface{}) interface{}) interface{} {
		if i < 0 { // Check out of bounds
			if i == -1 { // Check is identity
				return pair(func() (interface{}, interface{}) {
					return index, elements[0] // Return the index and the element
				})
			} else if i == -2 { // Check is list identity
				// Iterate through the provided callbacks
				for _, cb := range callback {
					// Run the callback
					if res := cb(index, list); res != nil {
						return res // Return the result of the callback
					}
				}

				return list(index, callback...) // Return the list
			} else if i == -3 && len(callback) == 1 { // Check is list set
				// Check got a response from the callback
				if res := callback[0](index, nil); res != nil {
					return makeList(index, list, append([]interface{}{res}, elements[1:]...)...) // Return the list
				}

				return list(i, callback...) // Do the callback for the other items
			}

			panic("index out of bounds") // Panic
		}

		// Iterate through the provided callbacks
		for _, cb := range callback {
			// Run the callback
			if res := cb(index, list); res != nil {
				return res // Return the result of the callback
			}
		}

		if i == index { // Check is node being addressed
			return elements[0] // Return the element
		}

		return list(i, callback...) // Return the element at that index elsewhere in the list
	}

	// Check is done constructing the list
	if len(elements) == 1 {
		return l // Return the final list
	}

	return makeList(index+1, l, elements[1:]...) // Return the list after the rest of the constructing calls
}

/* END INTERNAL METHODS */
