package shellsort

import (
	"fmt"
)

// ShellSort function implements shell sort algorithm
func ShellSort(arr []int) ([]int, int) {
	n := len(arr)
	iterations := 0

	// Start with a big gap, then reduce the gap
	for gap := n / 2; gap > 0; gap /= 2 {
		// do a gapped insertion sort for this gap size.
		// the first gap elements a[0..gap-1] are already
		// in gapped order keep adding one more element
		// until the entire array is gap sorted
		for i := gap; i < n; i++ {
			// add a[i] to the elements that have been gap
			// sorted save a[i] in temp and make a hole at
			// position i
			temp := arr[i]

			// shift earlier gap-sorted elements up until
			// the correct location for a[i] is found
			j := i
			for ; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
				iterations++
			}

			// put temp (the original a[i]) in its correct location
			arr[j] = temp
			iterations++
		}
	}
	return arr, iterations
}

// Test function tests ShellSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Shell sorting started\n")
	res, iterations := ShellSort(values[:])
	fmt.Printf("Shell sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Shell sorting result = %v\n", res)
}
