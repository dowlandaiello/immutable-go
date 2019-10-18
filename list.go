// Package immutable implements a set of functional, persistent, immutable data
// collections.
package immutable

// List is a list.
type List = ListNode

// Iterator if a generic method used to meaningfully traverse a list.
type Iterator = func(index int, value interface{}) interface{}

// ListNode is a node in a list.
type ListNode func(interface{}) interface{}

/* BEGIN EXPORTED METHODS */

// NewList initializes a new list from the given values.
func NewList(values ...interface{}) List {
	// Check list should be empty
	if len(values) == 0 {
		return nilListNode() // Return a nil list node
	}

	return newListNode(0, values[0], values[1:]...) // Return the list
}

// Size gets the size of the list.
func (l *List) Size() int {
	return (*l)(func(index int, value interface{}) interface{} {
		return index + 1 // Return the index
	}).(int) // Return the size of the list
}

// Push pushes an item to the list
func (l *List) Push(e interface{}) List {

}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// newListNode initializes a new list node from the given value and index in an
// overall list.
func newListNode(index int, value interface{}, children ...interface{}) ListNode {
	return func(f interface{}) interface{} {
		// Inputted parameter is an integer
		if i, ok := f.(int); ok {
			if i == index {
				return value // Return the value of the node
			}
		} else if i, ok := f.(Iterator); ok {
			if r := i(index, value); len(children) == 0 { // Check no more children
				return r // Return the result of the iterator
			}
		}

		var nextChildren []interface{} = nil // The children passed into the next iteration

		// Check no more children
		if len(children) == 0 {
			return nil // Stop execution
		} else if len(children) > 1 {
			nextChildren = children[1:] // Has next children
		}

		return newListNode(index+1, children[0], nextChildren...)(f) // Do for the rest of the nodes
	} // Return the list node
}

// nilListNode makes a list of size 0.
func nilListNode() ListNode {
	// A node of nil value with no children
	return func(f interface{}) interface{} {
		return nil // Return nil
	}
}

/* END INTERNAL METHODS */
