package mergesort

import (
	"fmt"
)

var iterations = 0

func merge(arr []int, l int, m int, r int) {
	// find arrays sizes
	lN := m - l + 1
	rN := r - m

	// create temp arrays
	lArr := make([]int, lN)
	rArr := make([]int, rN)

	// copy data to arrays
	for i := 0; i < lN; i++ {
		lArr[i] = arr[l+i]
		iterations++
	}
	for i := 0; i < rN; i++ {
		rArr[i] = arr[m+i+1]
		iterations++
	}

	// merge arrays
	i := 0
	j := 0
	k := l

	for i < lN && j < rN {
		if lArr[i] <= rArr[j] {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
		}
		k++
		iterations++
	}

	// copy remaining elements of lArr
	for i < lN {
		arr[k] = lArr[i]
		i++
		k++
		iterations++
	}

	// copy remaining elements of rArr
	for j < rN {
		arr[k] = rArr[j]
		j++
		k++
		iterations++
	}
}

// MergeSort function implements merge sort algorithm
func MergeSort(arr []int, l int, r int) {
	if r > l {
		middle := (l + r) / 2

		MergeSort(arr, l, middle)
		MergeSort(arr, middle+1, r)

		merge(arr, l, middle, r)
	}
}

// Test function tests MergeSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Merge sorting started\n")
	MergeSort(values[:], 0, len(values)-1)
	fmt.Printf("Merge sorting finished. Elements count = %v; Iterations taken = %v\n", len(values), iterations)
	fmt.Printf("Merge sorting result = %v\n", values)
}
