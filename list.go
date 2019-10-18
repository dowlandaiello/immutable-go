// Package immutable implements a set of functional, persistent, immutable data
// collections.
package immutable

// List is a list.
type List = ListNode

// ListNode is a node in a list.
type ListNode func(index int) interface{}

/* BEGIN EXPORTED METHODS */

// NewList initializes a new list from the given values.
func NewList(values ...interface{}) List {
	// Check list should be empty
	if len(values) == 0 {
		return nilListNode() // Return a nil list node
	}

	return newListNode(0, values[0], values[1:]...) // Return the list
}

// NewListFromSlice initializes a new list from the given values.
func NewListFromSlice(values []interface{}) List {
	return NewList(values...) // Return the new list
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// newListNode initializes a new list node from the given value and index in an
// overall list.
func newListNode(index int, value interface{}, children ...interface{}) ListNode {
	return func(i int) interface{} {
		// Check is addressed node
		if index == i {
			return value // Return the value
		}

		return newListNode(index+1, children[0], children[1:]...)(i) // Return the following nodes
	} // Return the list node
}

// nilListNode initializes a zero-value list node.
func nilListNode() ListNode {
	return func(i int) interface{} {
		return nil // Always nil
	}
}

/* END INTERNAL METHODS */
