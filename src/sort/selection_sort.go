package sort

import (
	"fmt"
)

// SelectionSort function implements selection sort algorithm
func SelectionSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n-i; j++ {
			iterations++
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr, iterations
}

// TestSelectionSort function tests SelectionSort
// You could provide elements array to sort
func TestSelectionSort(values []int) {
	fmt.Printf("Selection sorting started\n")
	res, iterations := SelectionSort(values[:])
	fmt.Printf("Selection sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Selection sorting result = %v\n", res)
}
