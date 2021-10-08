package gnomesort

import (
	"fmt"
)

var iterations = 0

// GnomeSort function implements gnome sort algorithm
func GnomeSort(arr []int) ([]int, int) {
	iterations := 0
	n := len(arr)

	i := 1
	for i != n-1 {
		if i == 0 {
			i++
		} else if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		} else {
			i++
		}
		iterations++
	}
	return arr, iterations
}

// Test function tests GnomeSort
// You could provide elements array to sort
func Test(values []int) {
	fmt.Printf("Gnome sorting started\n")
	res, iterations := GnomeSort(values[:])
	fmt.Printf("Gnome sorting finished. Elements count = %v; Iterations taken = %v\n", len(res), iterations)
	fmt.Printf("Gnome sorting result = %v\n", res)
}
