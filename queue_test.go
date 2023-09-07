package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// Initialize queue
	// Dequeue <- () - () - () <- Enqueue
	q := list.New()

	// Enqueue O(1)
	q.PushBack(1)
	q.PushBack(2)

	// Dequeue O(1)
	v, _ := q.Front().Value, q.Remove(q.Front())
	fmt.Println(v)

	// Peek O(1)
	v = q.Front().Value
	fmt.Println(v)

	// IsEmpty O(1)
	fmt.Println(q.Len() == 0)

	// Iterate over items O(N)
	fmt.Println("Iterate over all items")
	cur := q.Front()
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Next()
	}
}
