// Package immutable implements a set of functional, persistent, immutable data
// collections.
package immutable

import (
	"testing"
)

// testData is a data structure holding testing values.
type testData struct {
	Value int // the test value
}

/* BEGIN EXPORTED METHODS TESTS */

// TestNewList tests the functionality of the NewList helper method.
func TestNewList(t *testing.T) {
	list := NewList("des", "pa", "cito")                          // Put 3 elements in a new list
	if list(0) != "des" || list(1) != "pa" || list(2) != "cito" { // Check first element in the list is not equivalent to the generated testing integer
		t.Fatal("invalid list data") // Panic
	}
}

// TestGet tests the functionality of the NewList get functionality.
func TestGet(t *testing.T) {
	list := NewList(69, 420, 666)                          // Put 2 elements in a new list
	if list(0) != 69 || list(1) != 420 || list(2) != 666 { // Check last element in the list is not equivalent to the actual last item
		t.Fatal("invalid list data") // Panic
	}
}

// TestSize tests the Size list helper method.
func TestSize(t *testing.T) {
	list := NewList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9) // Make a new list with 10 integers
	if i := list.Size(); i != 10 {                // Check more than 10 integers in the list
		t.Fatalf("invalid list size; expected %d, got %d", 10, i) // Panic
	}
}

/* END EXPORTED METHODS TESTS */
