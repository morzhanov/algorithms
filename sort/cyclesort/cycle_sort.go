package cyclesort

import (
	"fmt"
)

// CycleSort function implements cycle sort algorithm
func CycleSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	// traverse array elements and put it to on
	// the right place
	for cycleStart := 0; cycleStart <= n-2; cycleStart++ {
		// initialize item as starting point
		item := arr[cycleStart]

		// find position where we put the item. We basically
		// count all smaller elements on right side of item.
		pos := cycleStart
		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
			iterations++
		}

		// if item is already in correct position
		if pos == cycleStart {
			continue
		}
		// ignore all duplicate  elements
		for item == arr[pos] {
			pos++
			iterations++
		}
		// put the item to it's right position
		if pos != cycleStart {
			item, arr[pos] = arr[pos], item
		}
		// rotate rest of the cycle
		for pos != cycleStart {
			pos = cycleStart
			// find position where we put the element
			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
				iterations++
			}
			// ignore all duplicate  elements
			for item == arr[pos] {
				pos++
				iterations++
			}
			// put the item to it's right position
			if item != arr[pos] {
				item, arr[pos] = arr[pos], item
			}
		}
		iterations++
	}
	return arr, iterations
}

// Test function tests CycleSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Cycle sorting cycleStarted\n")
	res, iterations := CycleSort(values[:])
	fmt.Printf("Cycle sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Cycle sorting result = %v\n", res)
}
