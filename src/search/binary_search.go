package search

import (
	"fmt"
)

// BinarySearch function implements binary search algorithm
// returns position of element and how much iterations used
func BinarySearch(arr []int, l int, r int, val int, iteration int) (int, int) {
	iteration++

	if r >= l {
		mid := ((r - l) / 2) + l

		if arr[mid] == val {
			return mid, iteration
		}
		if arr[mid] > val {
			return BinarySearch(arr, l, mid-1, val, iteration)
		}
		return BinarySearch(arr, mid+1, r, val, iteration)
	}
	return -1, iteration
}

// TestBinarySearch function tests BinarySearch
func TestBinarySearch(values []int, value int) {
	fmt.Printf("Binary search started, searching for element %v\n", value)
	pos, iterations := BinarySearch(values[:], 0, len(values)-1, value, 0)
	fmt.Printf("Binary search finished. Elements count = %v; Iterations taken = %v\n", len(values), iterations)
	fmt.Printf("Binary search elements: %v\n", values)
	if pos == -1 {
		fmt.Printf("Binary search result: element not found")
	} else {
		fmt.Printf("Binary search result: element found at position = %v\n", pos)
	}
}
