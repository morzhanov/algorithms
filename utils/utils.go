package utils

import (
	"math/rand"

	"github.com/morzhanov/algorithms/sort/quicksort"
)

// AlgorithmEnableFlags is a map with boolean values which disable/enable algorithm tests
type AlgorithmEnableFlags map[string]bool

// PrepareValuesForSeach creates array and value for search
func PrepareValuesForSeach(performSorting bool) ([]int, int) {
	values := make([]int, 100)
	value := 112
	for i := range values {
		values[i] = rand.Intn(500)
	}
	values[9] = value

	if performSorting {
		values = quicksort.QuickSort(values[:], 0, len(values)-1)
	}
	return values, value
}

// PrepareValuesForSort creates array for sort
func PrepareValuesForSort(count int) []int {
	values := make([]int, count)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

// PrepareFloatValuesForSort creates float array for sort
func PrepareFloatValuesForSort(count int) []float64 {
	values := make([]float64, count)
	for i := range values {
		values[i] = float64(rand.Intn(100)) / 100
	}
	return values
}

// IntComparator function compares int values
func IntComparator(a int, b int) int {
	if a > b {
		return 1
	}
	if b > a {
		return -1
	}
	return 0
}

// IntMax function returns max int
func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// IntMin function returns min int
func IntMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// IntMinArr function returns min int
func IntMinArr(ints []int) int {
	min := ints[0]
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}

// RuneMax function returns max rune
func RuneMax(a rune, b rune) rune {
	if a > b {
		return a
	}
	return b
}

// RuneMin function returns min rune
func RuneMin(a rune, b rune) rune {
	if a < b {
		return a
	}
	return b
}

// RuneMinArr function returns min rune
func RuneMinArr(runes []rune) rune {
	min := runes[0]
	for _, r := range runes {
		if r < min {
			min = r
		}
	}
	return min
}
