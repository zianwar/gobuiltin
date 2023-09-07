package main

import (
	"fmt"
	"testing"
)

func TestHashmap(t *testing.T) {
	// Initialize hashmap
	m := map[int]bool{}

	// Insert O(1)
	m[1] = true
	fmt.Println("Insert 1")
	m[2] = true
	fmt.Println("Insert 2")

	// Delete O(1)
	delete(m, 1)
	fmt.Println("Delete 1")

	// Lookup O(1)
	_, ok := m[1]
	fmt.Printf("Lookup 1 -> %v\n", ok)

	// IsEmpty O(1)
	fmt.Printf("IsEmpty -> %v\n", len(m) == 0)

	// Iterate over items O(N)
	fmt.Println("Iterate over all items")
	for v := range m {
		fmt.Println(v)
	}
}
