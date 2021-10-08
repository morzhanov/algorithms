package stupidsort

import (
	"fmt"
	"math/rand"
	"time"
)

// StupidSort function implements stupid sort algorithm
func StupidSort(arr []int) ([]int, int) {
	n := len(arr)
	iterations := 0
	for isSorted(arr) == false {
		iterations++
		for i := range arr {
			rand.Seed(time.Now().UnixNano())
			r := rand.Intn(n)
			arr[i], arr[r] = arr[r], arr[i]
		}
	}
	return arr, iterations
}

func isSorted(arr []int) bool {
	for i, v := range arr {
		if i != len(arr)-1 && v > arr[i+1] {
			return false
		}
	}
	return true
}

// Test function tests StupidSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Stupid sorting started\n")
	res, iterations := StupidSort(values[:])
	fmt.Printf("Stupid sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Stupid sorting result = %v\n", res)
}
