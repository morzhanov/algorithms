package search

import (
	"fmt"
	"math"
)

// JumpSearch function implements jump search algorithm
// returns position of element and how much iterations used
func JumpSearch(arr []int, val int) (int, int) {
	iterations := 0
	n := float64(len(arr))
	step := math.Floor(math.Sqrt(n))
	var prev = 0.0

	for arr[int(math.Min(step, n))-1] < val {
		prev = step
		step += math.Floor(math.Sqrt(n))
		if prev >= n {
			return -1, iterations
		}
	}

	pos, iteration := LinearSearch(arr[int(prev):], val)
	return pos, iteration + iterations
}

// TestJumpSearch function tests JumpSearch
func TestJumpSearch(values []int, value int) {
	fmt.Printf("Jump search started, searching for element %v\n", value)
	pos, iterations := JumpSearch(values[:], value)
	fmt.Printf("Jump search finished. Elements count = %v; Iterations taken = %v\n", len(values), iterations)
	fmt.Printf("Jump search elements: %v\n", values)
	if pos == -1 {
		fmt.Printf("Jump search result: element not found")
	} else {
		fmt.Printf("Jump search result: element found at position = %v\n", pos)
	}
}
