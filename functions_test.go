package main

import (
	"fmt"
	"testing"
)

func appendToSlice(a *[]int) {
	*a = append(*a, 1)
	fmt.Print("item added", (*a)[0])
}

// Closures
func outer() (func() int, int) {
	outerVar := 2
	inner := func() int {
		outerVar += 99 // outerVar from outer scope is mutated.
		return outerVar
	}
	inner()
	return inner, outerVar // return inner func and mutated outerVar 101
}

func TestFunctions(t *testing.T) {
	a := []int{}
	appendToSlice(&a)
	fmt.Println(a)
}
