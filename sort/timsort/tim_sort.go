package timsort

import (
	"fmt"

	"github.com/morzhanov/algorithms/utils"
)

var iterations = 0

const run = 32

// this function sorts array from left index to
// to right index which is of size atmost RUN
func insertionSort(arr []int, left int, right int) []int {
	for i := left + 1; i <= right; i++ {
		temp := arr[i]
		j := i - 1
		for j >= left && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
			iterations++
		}
		arr[j+1] = temp
		iterations++
	}
	return arr
}

// merge function merges the sorted runs
func merge(arr []int, l int, m int, r int) []int {
	// original array is broken in two parts
	// left and right array
	len1 := m - l + 1
	len2 := r - m
	left := make([]int, len1)
	right := make([]int, len2)
	for x := 0; x < len1; x++ {
		left[x] = arr[l+x]
		iterations++
	}
	for x := 0; x < len2; x++ {
		right[x] = arr[m+1+x]
		iterations++
	}

	i := 0
	j := 0
	k := l
	// after comparing, we merge those two array
	// in larger sub array
	for i < len1 && j < len2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
		iterations++
	}

	// copy remaining elements of left, if any
	for i < len1 {
		arr[k] = left[i]
		k++
		i++
		iterations++
	}
	// copy remaining element of right, if any
	for j < len2 {
		arr[k] = right[j]
		k++
		j++
		iterations++
	}
	return arr
}

// TimSort function implements tim sort algorithm
func TimSort(arr []int) []int {
	n := len(arr)
	// sort individual subarrays of size RUN
	for i := 0; i < n; i += run {
		insertionSort(arr, i, utils.IntMin((i+run-1), (n-1)))
		iterations++
	}

	// start merging from size RUN (or 32). It will merge
	// to form size 64, then 128, 256 and so on ....
	for size := run; size < n; size = 2 * size {
		// pick starting point of left sub array. We
		// are going to merge arr[left..left+size-1]
		// and arr[left+size, left+2*size-1]
		// After every merge, we increase left by 2*size
		for left := 0; left < n; left += 2 * size {
			// find ending point of left sub array
			// mid+1 is starting point of right sub array
			mid := left + size - 1
			right := utils.IntMin((left + 2*size - 1), (n - 1))
			// merge sub array arr[left.....mid] &
			// arr[mid+1....right]
			merge(arr, left, mid, right)
			iterations++
		}
	}
	return arr
}

// Test function tests TimSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Tim sorting started\n")
	res := TimSort(values[:])
	fmt.Printf("Tim sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Tim sorting result = %v\n", res)
}
