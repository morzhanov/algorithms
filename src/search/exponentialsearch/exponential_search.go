package exponentialsearch

import (
	"fmt"

	"github.com/morzhanov/algorithms/src/search/binarysearch"
	"github.com/morzhanov/algorithms/src/utils"
)

// ExponentialSearch function implements interpolation search algorithm
// returns position of element and how much iterations used
func ExponentialSearch(arr []int, n int, val int, iteration int) (int, int) {
	iteration++
	if arr[0] == val {
		return 0, iteration
	}

	// find range for binary search by repeated doubling
	i := 1
	for i < n && arr[i] <= val {
		i *= 2
	}

	// call binary search for the found range.
	return binarysearch.BinarySearch(arr, i/2, utils.IntMin(i, n), val, iteration)
}

// Test function tests ExponentialSearch
func Test(values []int, value int) {
	fmt.Printf("Exponential search started, searching for element %v\n", value)
	pos, iterations := ExponentialSearch(values[:], len(values)-1, value, 0)
	fmt.Printf("Exponential search finished. Elements count = %v; Iterations taken = %v\n", len(values), iterations)
	fmt.Printf("Exponential search elements: %v\n", values)
	if pos == -1 {
		fmt.Printf("Exponential search result: element not found")
	} else {
		fmt.Printf("Exponential search result: element found at position = %v\n", pos)
	}
}
