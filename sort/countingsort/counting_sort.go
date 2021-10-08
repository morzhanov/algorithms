package countingsort

import (
	"fmt"
)

func findMinMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// CountingSort function implements counting sort algorithm
func CountingSort(arr []int) ([]int, int) {
	iterations := 0
	min, max := findMinMax(arr)
	count := make([]int, max-min+2)

	// count
	for _, e := range arr {
		count[e]++
		iterations++
	}

	// add prev val to curr
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
		iterations++
	}

	// copy to correct pos
	res := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		e := arr[i]
		t := count[e] - 1
		res[t] = e
		count[e] = count[e] - 1
		iterations++
	}

	return res, iterations
}

// Test function tests CountingSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Counting sorting started\n")
	res, iterations := CountingSort(values[:])
	fmt.Printf("Counting sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Counting sorting result = %v\n", res)
}
