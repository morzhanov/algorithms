package pairedinsertionsort

import (
	"fmt"
)

// PairedInsertionSort function implements paired insertion sort algorithm
func PairedInsertionSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	i := 0
	for ; i < n-1; i++ {
		// let x and y hold the next to elements in array
		x := arr[i]
		y := arr[i+1]

		// ensure that x is not smaller than y
		if x < y {
			x, y = y, x
		}

		// j is the index used to find the insertion point
		// find the insertion point for x
		// shift existing content by 2
		j := i - 1
		for ; j >= 0 && arr[j] > x; j-- {
			arr[j+2] = arr[j]
			iterations++
		}

		// store x at its insertion place
		// arr[j+1] is an available space now
		arr[j+2] = x

		//find the insertion point for y
		// shift existing content by 1
		for j >= 0 && arr[j] > y {
			arr[j+1] = arr[j]
			j = j - 1
			iterations++
		}
		// store y at its insertion place
		arr[j+1] = y
	}

	// if length(A) is odd, an extra
	// single insertion is needed for
	// the last element
	if i == len(arr)-1 {
		y := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > y; j-- {
			arr[j+1] = arr[j]
			iterations++
		}
		arr[j+1] = y
	}
	return arr, iterations
}

// Test function tests PairedInsertionSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Paired Insertion sorting started\n")
	res, iterations := PairedInsertionSort(values[:])
	fmt.Printf("Paired Insertion sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Paired Insertion sorting result = %v\n", res)
}
