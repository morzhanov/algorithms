package oddevensort

import (
	"fmt"
)

// OddEvenSort function implements odd-even sort algorithm
func OddEvenSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	isSorted := false
	for !isSorted {
		isSorted = true
		// perform bubble sort on odd indexed element
		for i := 1; i <= n-2; i = i + 2 {
			iterations++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		// perform bubble sort on even indexed element
		for i := 0; i <= n-2; i = i + 2 {
			iterations++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr, iterations
}

// Test function tests OddEvenSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("OddEven sorting started\n")
	res, iterations := OddEvenSort(values[:])
	fmt.Printf("OddEven sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("OddEven sorting result = %v\n", res)
}
