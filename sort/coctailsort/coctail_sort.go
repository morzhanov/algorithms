package coctailsort

import (
	"fmt"
)

// CoctailSort function implements coctail sort algorithm
func CoctailSort(arr []int) ([]int, int) {
	n := len(arr)
	iterations := 0
	swapped := true
	start := 0
	end := n - 1

	for swapped {
		swapped = false
		for i := start; i < end-1; i++ {
			iterations++
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			swapped = true
		}
		if !swapped {
			break
		}
		swapped = false
		end--
		for i := end - 1; i > start; i-- {
			iterations++
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
			swapped = true
		}
		if !swapped {
			break
		}
		start++
	}

	return arr, iterations
}

// Test function tests CoctailSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Coctail sorting started\n")
	res, iterations := CoctailSort(values[:])
	fmt.Printf("Coctail sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Coctail sorting result = %v\n", res)
}
