package stack

import (
	"fmt"
	"time"
)

// Stack implementation
// https://www.geeksforgeeks.org/stack-data-structure/
// Access: O(n)
// Search: O(n)
// Insert: O(1)
// Delete: O(1)
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

// Pop a last value from Stack, return true if stack is empty
func (s *Stack) Pop() (float64, bool) {
	if s.IsEmpty() {
		return 0, true
	}

	lastIndex := len(*s) - 1
	lastItem := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return lastItem, false
}

// Top returns a last value from Stack, return true if stack is empty
func (s *Stack) Top() (float64, bool) {
	if s.IsEmpty() {
		return 0, true
	}

	return (*s)[len(*s)-1], false
}

// Test is Stack implememntation test
func Test(s Stack) {
	fmt.Printf("Stack length: %v\n", s.Len())

	fmt.Printf("Pushing to stack %v\n", 10)
	s.Push(10)
	fmt.Printf("Pushing to stack %v\n", 20)
	s.Push(20)
	top, empty := s.Top()
	fmt.Printf("Top value = %v. Stack empty = %v\n", top, empty)
	fmt.Printf("Stack length = %v\n", s.Len())

	fmt.Printf("Pop value from stack\n")
	value, empty := s.Pop()
	fmt.Printf("Value = %v. Stack empty = %v\n", value, empty)
	fmt.Printf("Pop value from stack\n")
	value, empty = s.Pop()
	fmt.Printf("Value = %v. Stack empty = %v\n", value, empty)

	top, empty = s.Top()
	fmt.Printf("Top value = %v. Stack empty = %v\n", top, empty)
	fmt.Printf("Stack length = %v\n", s.Len())

	fmt.Printf("Adding 1K elements\n")
	before := time.Now()
	for i := 0.0; i < 1000; i++ {
		s.Push(i)
	}
	after := time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Extracting all 1K elements\n")
	before = time.Now()
	for i := 0.0; i < 1000; i++ {
		s.Pop()
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Adding 100M elements\n")
	before = time.Now()
	for i := 0.0; i < 100000000; i++ {
		s.Push(i)
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Extracting all 100M elements\n")
	before = time.Now()
	for i := 0.0; i < 100000000; i++ {
		s.Pop()
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n\n\n", after.Sub(before))
}
