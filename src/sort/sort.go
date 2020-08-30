package sort

import (
	"github.com/morzhanov/algorithms/src/sort/bubblesort"
	"github.com/morzhanov/algorithms/src/sort/countingsort"
	"github.com/morzhanov/algorithms/src/sort/insertionsort"
	"github.com/morzhanov/algorithms/src/sort/mergesort"
	"github.com/morzhanov/algorithms/src/sort/selectionsort"
	"github.com/morzhanov/algorithms/src/sort/stupidsort"
	"github.com/morzhanov/algorithms/src/utils"
)

// RunSortTests function runs sorting tests
func RunSortTests(flags utils.AlgorithmEnableFlags) {
	if flags["stupid"] == true {
		// StupidSorting
		// extremely stupid sorting,
		// do not provide more than 5-10 elements or app could stuck
		values := utils.PrepareValuesForSort(5)
		stupidsort.Test(values[:])
	}
	if flags["bubble"] == true {
		// BubbleSorting
		values := utils.PrepareValuesForSort(100)
		bubblesort.Test(values[:])
	}
	if flags["counting"] == true {
		// CountingSorting
		values := utils.PrepareValuesForSort(100)
		countingsort.Test(values[:])
	}
	if flags["selection"] == true {
		// SelectionSorting
		values := utils.PrepareValuesForSort(100)
		selectionsort.Test(values[:])
	}
	if flags["merge"] == true {
		// MergeSorting
		values := utils.PrepareValuesForSort(100)
		mergesort.Test(values[:])
	}
	if flags["insert"] == true {
		// InsertionSorting
		values := utils.PrepareValuesForSort(100)
		insertionsort.Test(values[:])
	}
}
