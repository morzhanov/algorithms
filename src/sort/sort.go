package sort

import (
	"github.com/morzhanov/algorithms/src/sort/bubblesort"
	"github.com/morzhanov/algorithms/src/sort/coctailsort"
	"github.com/morzhanov/algorithms/src/sort/combsort"
	"github.com/morzhanov/algorithms/src/sort/countingsort"
	"github.com/morzhanov/algorithms/src/sort/insertionsort"
	"github.com/morzhanov/algorithms/src/sort/mergesort"
	"github.com/morzhanov/algorithms/src/sort/oddevensort"
	"github.com/morzhanov/algorithms/src/sort/pairedinsertionsort"
	"github.com/morzhanov/algorithms/src/sort/selectionsort"
	"github.com/morzhanov/algorithms/src/sort/shellsort"
	"github.com/morzhanov/algorithms/src/sort/stupidsort"
	"github.com/morzhanov/algorithms/src/sort/treesort"
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
	if flags["pairedInsert"] == true {
		// PairedInsertionSorting
		values := utils.PrepareValuesForSort(100)
		pairedinsertionsort.Test(values[:])
	}
	if flags["pairedInsert"] == true {
		// PairedInsertionSorting
		values := utils.PrepareValuesForSort(100)
		pairedinsertionsort.Test(values[:])
	}
	if flags["shell"] == true {
		// ShellSorting
		values := utils.PrepareValuesForSort(100)
		shellsort.Test(values[:])
	}
	if flags["coctail"] == true {
		// CoctailSorting
		values := utils.PrepareValuesForSort(100)
		coctailsort.Test(values[:])
	}
	if flags["tree"] == true {
		// TreeSorting
		values := utils.PrepareValuesForSort(100)
		treesort.Test(values[:])
	}
	if flags["oddEven"] == true {
		// OddEvenSorting
		values := utils.PrepareValuesForSort(100)
		oddevensort.Test(values[:])
	}
	if flags["comb"] == true {
		// CombSorting
		values := utils.PrepareValuesForSort(100)
		combsort.Test(values[:])
	}
}
