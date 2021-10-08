package combsort

import (
	"fmt"
)

// CombSort function implements comb sort algorithm
func CombSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	coef := 1.3
	gap := int(float64(n) / coef)
	for ; gap != 0; gap = int(float64(gap) / coef) {
		for i := 0; i+gap < n-1; i++ {
			iterations++
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}

	return arr, iterations
}

// Test function tests CombSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Comb sorting started\n")
	res, iterations := CombSort(values[:])
	fmt.Printf("Comb sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Comb sorting result = %v\n", res)
}
