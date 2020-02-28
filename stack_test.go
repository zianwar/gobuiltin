package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var v interface{}

	// Initialize stack
	s := list.New()

	// Push O(1)
	s.PushBack(1)
	s.PushBack(2)

	// Pop O(1)
	v, _ = s.Back().Value, s.Remove(s.Back())
	fmt.Println(v)

	// Peek O(1)
	v = s.Back().Value
	fmt.Println(v)

	// IsEmpty O(1)
	fmt.Println(s.Len() == 0)

	// Iterate over items O(N)
	cur := s.Front()
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Next()
	}
}
