package ds

// Stack implementation
// https://www.geeksforgeeks.org/stack-data-structure/
type Stack []float64

// IsEmpty returns is Stack empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns Stack length
func (s *Stack) Len() int {
	return len(*s)
}

// Push a new value to Stack
func (s *Stack) Push(item float64) {
	*s = append(*s, item)
}

// Pop a last value from Stack, return false if stack is empty
func (s *Stack) Pop() (float64, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	lastIndex := len(*s) - 1
	lastItem := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return lastItem, true
}
