package linearsearch

import (
	"fmt"
)

// LinearSearch function implements linear search algorithm
// returns position of element and how much iterations used
func LinearSearch(arr []int, val int) (int, int) {
	iterations := 0
	for i, v := range arr {
		iterations++
		if v == val {
			return i, iterations
		}
	}
	return -1, iterations
}

// Test function tests LinearSearch
func Test(values []int, value int) {
	fmt.Printf("Linear search started, searching for element %v\n", value)
	pos, iterations := LinearSearch(values[:], value)
	fmt.Printf("Linear search finished. Elements count = %v; Iterations taken = %v\n", len(values), iterations)
	fmt.Printf("Linear search elements: %v\n", values)
	if pos == -1 {
		fmt.Printf("Linear search result: element not found")
	} else {
		fmt.Printf("Linear search result: element found at position = %v\n", pos)
	}
}
