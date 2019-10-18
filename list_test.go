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

/* END EXPORTED METHODS TESTS */

/* BEGIN INTERNAL METHODS */

// generateTestData generates a slice of test strings.
func generateTestData(n int) []testData {
	var t []testData // The test data

	// Make n ints
	for i := 0; i < n; i++ {
		t = append(t, testData{rand.Intn(69)}) // Add the test integer to the test data slice
	}

	return t // Return the test data
}

/* END INTERNAL METHODS */
