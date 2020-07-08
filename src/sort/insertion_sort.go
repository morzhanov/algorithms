package sort

import (
	"fmt"
)

// InsertionSort function implements counting sort algorithm
func InsertionSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	for i := 1; i < n-1; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
			iterations++
		}
		arr[j+1] = key
		iterations++
	}

	return arr, iterations
}

// TestInsertionSort function tests InsertionSort
// You could provide elements array to sort
func TestInsertionSort(values []int) {
	fmt.Printf("Insertion sorting started\n")
	res, iterations := InsertionSort(values[:])
	fmt.Printf("Insertion sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Insertion sorting result = %v\n", res)
}
