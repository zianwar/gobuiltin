package main

import (
	"container/heap"
	"fmt"
	"testing"
)

type Heap struct {
	items []int
}

func (h *Heap) Len() int           { return len(h.items) }
func (h *Heap) Less(i, j int) bool { return h.items[i] < h.items[j] }
func (h *Heap) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *Heap) Push(v interface{}) { h.items = append(h.items, v.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &Heap{items: make([]int, 0)}
	heap.Init(h)

	heap.Push(h, 1)
	heap.Push(h, -1)
	heap.Push(h, 10)

	for h.Len() > 0 {
		x := heap.Pop(h).(int)
		fmt.Println(x)
	}
}
