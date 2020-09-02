package stoogesort

import (
	"fmt"
	"math"
)

var iterations = 0

// StoogeSort function implements stooge sort algorithm
func StoogeSort(arr []int, start int, end int) []int {
	if start >= end {
		return arr
	}
	iterations++
	if arr[start] > arr[end] {
		arr[start], arr[end] = arr[end], arr[start]
	}

	if end-start+1 > 2 {
		third := int(math.Ceil(float64(end-start-1) / 3))
		arr = StoogeSort(arr, start, end-third)
		arr = StoogeSort(arr, third+start, end)
		arr = StoogeSort(arr, start, end-third)
	}
	return arr
}

// Test function tests StoogeSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Stooge sorting started\n")
	res := StoogeSort(values[:], 0, len(values)-1)
	fmt.Printf("Stooge sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Stooge sorting result = %v\n", res)
}
