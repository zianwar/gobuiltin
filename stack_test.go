package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	// Initialize stack
	// () - () - () <-> Push/Pop
	s := list.New()

	// Push O(1)
	s.PushBack(1)
	fmt.Printf("Push 1\n")
	s.PushBack(2)
	fmt.Printf("Push 2\n")

	// Pop O(1)
	v, _ := s.Back().Value, s.Remove(s.Back())
	fmt.Printf("Pop -> %v\n", v)

	// Peek O(1)
	v = s.Back().Value
	fmt.Printf("Peek -> %v\n", v)

	// IsEmpty O(1)
	fmt.Printf("IsEmpty -> %v\n", s.Len() == 0)

	// Iterate over items O(N)
	fmt.Println("Iterate over all items")
	cur := s.Front()
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Next()
	}
}
