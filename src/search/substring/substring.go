package substring

import (
	"fmt"
)

// Substring function returns index of substring
func Substring(substr string, str string) {
	substrArr := []rune(substr)
	strArr := []rune(str)

	for i := 0; i <= len(strArr)-1; i++ {
		if len(substrArr)+i-1 > len(strArr)+1 {
			break
		}
		found := true
		for j := 0; j <= len(substrArr)-1; j++ {
			if strArr[i+j] != substrArr[j] {
				found = false
				break
			}
		}
		if found {
			fmt.Printf("Substring %v found in string %v at index %v\n", substr, str, i)
			return
		}
	}
	fmt.Printf("Substring %v is not found in string %v\n", substr, str)
}

// Test function tests Substring function
func Test(A string, B string) {
	Substring(A, B)
}
