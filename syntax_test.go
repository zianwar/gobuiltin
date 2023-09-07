package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSyntax(t *testing.T) {
	// Stack implementation
	l := list.New()

	// Push O(1)
	l.PushFront(1)
	println("Push 1")
	l.PushFront(2)
	println("Push 2")

	// Pop O(1)
	v, _ := l.Front().Value, l.Remove(l.Front())
	fmt.Printf("Pop -> %v\n", v)

	// Peek O(1)
	v = l.Front().Value
	fmt.Printf("Peek -> %v\n", v)

	// IsEmpty O(1)
	isEmpty := l.Len() == 0
	fmt.Printf("IsEmpty -> %v\n", isEmpty)
}
