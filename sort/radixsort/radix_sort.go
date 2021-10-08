package radixsort

import (
	"fmt"
)

func findMax(arr []int) (idx int, res int) {
	for i, val := range arr {
		if val > res {
			res = val
			idx = i
		}
	}
	return idx, res
}

func countSort(arr []int, n int, exp int) (res []int, iterations int) {
	res = make([]int, n)
	i := 0
	count := make([]int, 10)

	// store count of occurrences in count[]
	for i = 0; i < n; i++ {
		iterations++
		count[(arr[i]/exp)%10]++
	}

	// change count[i] so that count[i] now contains
	// actual position of this digit in output[]
	for i = 1; i < 10; i++ {
		iterations++
		count[i] += count[i-1]
	}

	// build the output array
	for i = n - 1; i >= 0; i-- {
		iterations++
		res[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	// copy the output array to arr[], so that arr[] now
	// contains sorted numbers according to curent digit
	for i = 0; i < n; i++ {
		iterations++
		arr[i] = res[i]
	}
	return res, iterations
}

// RadixSort function implements radix sort algorithm
func RadixSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)
	_, max := findMax(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		res, iter := countSort(arr, n, exp)
		arr = res
		iterations++
		iterations += iter
	}

	return arr, iterations
}

// Test function tests RadixSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Radix sorting started\n")
	res, iterations := RadixSort(values[:])
	fmt.Printf("Radix sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Radix sorting result = %v\n", res)
}
