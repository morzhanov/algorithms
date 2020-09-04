package interpolationserach

import (
	"fmt"
)

// InterpolationSearch function implements interpolation search algorithm
// returns position of element and how much iterations used
func InterpolationSearch(arr []int, l int, h int, val int, iteration int) (int, int) {
	iteration++
	// since array is sorted, an element
	// present in array must be in range
	// defined by corner
	if l <= h && val >= arr[l] && val <= arr[h] {
		// probing the position with keeping
		// uniform distribution in mind.
		pos := l + (((h - l) / (arr[h] - arr[l])) * (val - arr[l]))

		// condition of target found
		if arr[pos] == val {
			return pos, iteration
		}

		// if val is larger, val is in right sub array
		if arr[pos] < val {
			return InterpolationSearch(arr, pos+1, h, val, iteration)
		}

		// if val is smaller, val is in left sub array
		if arr[pos] > val {
			return InterpolationSearch(arr, l, pos-1, val, iteration)
		}
	}
	return -1, iteration
}

// Test function tests InterpolationSearch
func Test(values []int, value int) {
	fmt.Printf("Interpolation search started, searching for element %v\n", value)
	pos, iterations := InterpolationSearch(values[:], 0, len(values)-1, value, 0)
	fmt.Printf("Interpolation search finished. Elements count = %v; Iterations taken = %v\n", len(values), iterations)
	fmt.Printf("Interpolation search elements: %v\n", values)
	if pos == -1 {
		fmt.Printf("Interpolation search result: element not found")
	} else {
		fmt.Printf("Interpolation search result: element found at position = %v\n", pos)
	}
}
