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
	fmt.Println("Add 1")
	set[2] = true
	fmt.Println("Add 2")

	// Delete O(1)
	delete(set, 1)
	fmt.Println("Delete 1")

	// Contains O(1)
	_, ok := set[1]
	fmt.Printf("Contains 1 -> %v\n", ok)

	// IsEmpty O(1)
	fmt.Printf("IsEmpty -> %v\n", len(set) == 0)

	// Iterate over items O(N)
	fmt.Println("Iterate over all items")
	for v := range set {
		fmt.Println(v)
	}
}
