package main

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	// Initialize set
	set := map[int]bool{}

	// Add O(1)
	set[1] = true
	set[2] = true

	// Delete O(1)
	delete(set, 1)

	// Contains O(1)
	_, ok := set[1]
	fmt.Println(ok)

	// IsEmpty O(1)
	fmt.Println(len(set) == 0)

	// Iterate over items O(N)
	for v := range set {
		fmt.Println(v)
	}
}
