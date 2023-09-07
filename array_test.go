package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// Initialize hashmap
	a := []int{}

	// Append O(1)
	a = append(a, 1)
	fmt.Println("Append 1")

	// Append O(1)
	a = append(a, 2)
	fmt.Println("Append 2")

	// Lookup O(1)
	fmt.Printf("Lookup a[0] -> %v\n", a[0])

	// Insert O(n)
	// [1, 2] -> [1] + 3 + [2] -> [1, 3, 2]
	a = append(a[:1], append([]int{3}, a[1:]...)...)
	fmt.Printf("Insert 3 at index 1 -> %v\n", a)

	// IsEmpty O(1)
	fmt.Printf("IsEmpty -> %v\n", len(a) == 0)

	// Iterate over items O(N)
	fmt.Println("Iterate over all items")
	for _, v := range a {
		fmt.Println(v)
	}
}
