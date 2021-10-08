package sort

import (
	"github.com/morzhanov/algorithms/sort/bubblesort"
	"github.com/morzhanov/algorithms/sort/bucketsort"
	"github.com/morzhanov/algorithms/sort/coctailsort"
	"github.com/morzhanov/algorithms/sort/combsort"
	"github.com/morzhanov/algorithms/sort/countingsort"
	"github.com/morzhanov/algorithms/sort/cyclesort"
	"github.com/morzhanov/algorithms/sort/gnomesort"
	"github.com/morzhanov/algorithms/sort/insertionsort"
	"github.com/morzhanov/algorithms/sort/mergesort"
	"github.com/morzhanov/algorithms/sort/oddevensort"
	"github.com/morzhanov/algorithms/sort/pairedinsertionsort"
	"github.com/morzhanov/algorithms/sort/pigeonholesort"
	"github.com/morzhanov/algorithms/sort/quicksort"
	"github.com/morzhanov/algorithms/sort/radixsort"
	"github.com/morzhanov/algorithms/sort/selectionsort"
	"github.com/morzhanov/algorithms/sort/shellsort"
	"github.com/morzhanov/algorithms/sort/stoogesort"
	"github.com/morzhanov/algorithms/sort/stupidsort"
	"github.com/morzhanov/algorithms/sort/timsort"
	"github.com/morzhanov/algorithms/sort/treesort"
	"github.com/morzhanov/algorithms/utils"
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
	if flags["radix"] == true {
		// RadixSorting
		values := utils.PrepareValuesForSort(100)
		radixsort.Test(values[:])
	}
	if flags["pigeonhole"] == true {
		// PigeonholeSorting
		values := utils.PrepareValuesForSort(100)
		pigeonholesort.Test(values[:])
	}
	if flags["stooge"] == true {
		// StoogeSorting
		values := utils.PrepareValuesForSort(100)
		stoogesort.Test(values[:])
	}
	if flags["gnome"] == true {
		// GnomeSorting
		values := utils.PrepareValuesForSort(100)
		gnomesort.Test(values[:])
	}
	if flags["bucket"] == true {
		// BucketSorting
		values := utils.PrepareFloatValuesForSort(100)
		bucketsort.Test(values[:])
	}
	if flags["tim"] == true {
		// TimSorting
		values := utils.PrepareValuesForSort(100)
		timsort.Test(values[:])
	}
	if flags["cycle"] == true {
		// CycleSorting
		values := utils.PrepareValuesForSort(100)
		cyclesort.Test(values[:])
	}
	if flags["quick"] == true {
		// QuickSorting
		values := utils.PrepareValuesForSort(100)
		quicksort.Test(values[:])
	}
}
