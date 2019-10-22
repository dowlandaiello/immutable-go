// Package immutable implements a set of functional, persistent, immutable data
// collections.
package immutable

import (
	"math/rand"
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
	list := NewList(0) // Put 1 element in a new list

	for i := 1; i < 10000000; i++ {
		list = list.Push(rand.Intn(420)) // Add a random int to the list
		list(i)                          // Get the item at the index
	}
}

// TestForEach tests the functionality of the ForEach helper method.
func TestForEach(t *testing.T) {
	i := 0 // The iterator

	list := NewList(1, 2, 3, 4) // Put 4 items in a new list
	list.ForEach(func(index int, value interface{}) interface{} {
		i++ // Increment i

		return nil // Nothing to return
	})

	// Should have been incremented 4 times
	if i != 4 {
		t.Fatalf("callback should have been run 4 times; ran %d times instead", i) // Panic
	}
}

// TestSet tests the functionality of the Set helper method.
func TestSet(t *testing.T) {
	list := NewList(1, 2, 3, 4) // Put 4 items in a new list
	l1 := list.Set(2, 5)        // Set the 3rd index to 5

	if l1(0) != 1 || l1(1) != 2 || l1(2) != 5 || l1(3) != 4 { // Ensure item at 3rd index was set to 5
		t.Fatal("item should have been set") // Panic
	}
}

// TestPush tests the functionality of the Push helper method.
func TestPush(t *testing.T) {
	list := NewList(1, 2, 3)  // Put 3 elements in a new list
	if list.Push(4)(3) != 4 { // Check last element not added to list
		t.Fatal("element should have been added to list") // Panic
	}
}

// TestPop tests the functionality of the Pop helper method.
func TestPop(t *testing.T) {
	list := NewList(1, 2, 3, 4) // Put 4 elements in a new list
	if list.Pop().Size() != 3 { // Check last element not added to list
		t.Fatal("element should have been popped from list") // Panic
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
