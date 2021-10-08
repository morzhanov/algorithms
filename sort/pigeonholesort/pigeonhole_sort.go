package pigeonholesort

import (
	"fmt"
)

// PigeonholeSort function implements pigeonhole sort algorithm
func PigeonholeSort(arr []int) ([]int, int) {
	iterations := 0
	max := arr[0]
	min := arr[0]
	for _, val := range arr {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}

	holes := make([]int, max-min+1)
	for i := 0; i < len(arr); i++ {
		iterations++
		holes[arr[i]-min]++
	}
	idx := 0
	for i := 0; i < len(holes)-1; i++ {
		iterations++
		for holes[i] != 0 {
			iterations++
			arr[idx] = i + min
			idx++
			holes[i]--
		}
	}

	return arr, iterations
}

// Test function tests PigeonholeSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Pigeonhole sorting started\n")
	res, iterations := PigeonholeSort(values[:])
	fmt.Printf("Pigeonhole sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Pigeonhole sorting result = %v\n", res)
}
