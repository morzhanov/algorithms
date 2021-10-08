package bucketsort

import (
	"fmt"
	"math"
)

func insertionSort(arr []float64) ([]float64, int) {
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

// BucketSort function implements bucket sort algorithm
func BucketSort(arr []float64) ([]float64, int) {
	iterations := 0
	n := len(arr)
	buckets := make([][]float64, n-1)
	for i := 0; i < n-1; i++ {
		iterations++
		buckets[i] = make([]float64, 0)
	}

	// put elements into buckets
	for i := 0; i < n-1; i++ {
		iterations++
		curr := int(math.Floor(float64(arr[i]) * 10))
		buckets[curr] = append(buckets[curr], arr[i])
	}
	// sort each bucket with insertion sort
	for i := 0; i < n-1; i++ {
		iter := 0
		buckets[i], iter = insertionSort(buckets[i])
		iterations += iter
	}

	// fill origin arry
	idx := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			iterations++
			arr[idx] = buckets[i][j]
			idx++
		}
	}
	return arr, iterations
}

// Test function tests BucketSort
// You could provide elements array to sort
func Test(values []float64) {
	fmt.Printf("Bucket sorting started\n")
	res, iterations := BucketSort(values[:])
	fmt.Printf("Bucket sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Bucket sorting result = %v\n", res)
}
