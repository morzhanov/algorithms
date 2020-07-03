package sort

import (
	"fmt"
)

// BubbleSort function implements bubble sort algorithm
func BubbleSort(arr []int) ([]int, int) {
	iterations := 0
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			iterations++
			if arr[j-1] > arr[j] {
				intermediate := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = intermediate
			}
		}
	}
	return arr, iterations
}

// BubbleSortAlt function implements alternative bubble sort algorithm
func BubbleSortAlt(arr []int) ([]int, int) {
	var (
		n          = len(arr)
		sorted     = false
		iterations = 0
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			iterations++
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		sorted = !swapped
		n = n - 1
	}
	return arr, iterations
}

// TestBubbleSort function tests StupidSort
// You could provide elements array to sort
func TestBubbleSort(values []int) {
	fmt.Printf("Bubble sorting started\n")
	res, iterations := BubbleSort(values[:])
	fmt.Printf("Bubble sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Bubble sorting result = %v\n", res)

	fmt.Printf("Alternative Bubble sorting started (it will take less time if array partly sorted)\n")
	res2, iterations2 := BubbleSortAlt(values[:])
	fmt.Printf("Alternative Bubble sorting finished. Elements count = %v; Iterations taken = %v\n", len(res2), iterations2)
	fmt.Printf("Alternative Bubble sorting result = %v\n", res2)
}
