package main

import (
	"fmt"
	slicePlus "homework01/slice_plus"
)

func testRemove[T any](sample []T, index int) {
	fmt.Println("Before example1: ", sample)
	removedElement := slicePlus.SliceRemove(&sample, index)
	fmt.Println("After example1:", sample)
	fmt.Println("Removed element: ", removedElement)
	fmt.Println()
}

func main() {
	example1 := []int{1, 2, 3, 4, 5}
	testRemove(example1, 2)
	testRemove(example1, 3)

	example2 := []any{1, "123", 1.2}
	testRemove(example2, 0)
	testRemove(example2, 1)
	testRemove(example2, 2)
}
