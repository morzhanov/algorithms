package quicksort

import (
	"fmt"
)

var iterations = 0

// partition function takes last element as pivot,
// places the pivot element at its correct
// position in sorted array, and places all
// smaller (smaller than pivot) to left of
// pivot and all greater elements to right of pivot
func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := (low - 1) // index of smaller element
	for j := low; j < high; j++ {
		iterations++
		// if current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// swap arr[i+1] and arr[high] (or pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSort function implements quick sort algorithm
func QuickSort(arr []int, low int, high int) []int {
	if low < high {
		// pi is partitioning index, arr[pi] is now at right place
		pi := partition(arr, low, high)
		// recursively sort elements before partition and after partition
		arr = QuickSort(arr, low, pi-1)
		arr = QuickSort(arr, pi+1, high)
	}
	return arr
}

// Test function tests QuickSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Quick sorting cycleStarted\n")
	res := QuickSort(values[:], 0, len(values)-1)
	fmt.Printf("Quick sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Quick sorting result = %v\n", res)
}
