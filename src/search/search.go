package search

import (
	"github.com/morzhanov/algorithms/src/search/binarysearch"
	"github.com/morzhanov/algorithms/src/search/jumpsearch"
	"github.com/morzhanov/algorithms/src/search/linearsearch"
	"github.com/morzhanov/algorithms/src/utils"
)

// RunSearchTests function runs search tests
func RunSearchTests(flags utils.AlgorithmEnableFlags) {
	if flags["binary"] == true {
		// BinarySearch
		// NOTE: array should be sorted
		values, value := utils.PrepareValuesForSeach(true)
		binarysearch.Test(values[:], value)
	}
	if flags["linear"] == true {
		// LinearSearch
		values, value := utils.PrepareValuesForSeach(false)
		linearsearch.Test(values[:], value)
	}
	if flags["jump"] == true {
		// JumpSearch
		// NOTE: array should be sorted
		values, value := utils.PrepareValuesForSeach(false)
		jumpsearch.Test(values[:], value)
	}
}
