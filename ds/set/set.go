package set

import (
	"errors"
	"fmt"
)

// Set implementation
// https://blog.benoitvallon.com/data-structures-in-javascript/the-set-data-structure/
// Search: O(n)
// Insert: O(n)
// Delete: O(n)
type Set []float64

// Contains medhod returns true if element exists in Set
func (s *Set) Contains(item float64) bool {
	for _, v := range *s {
		if v == item {
			return true
		}
	}
	return false
}

// Add method is used to add value to Set
func (s *Set) Add(item float64) {
	if s.Contains(item) == false {
		*s = append(*s, item)
	}
}

// Delete method is used to delete value from Set
func (s *Set) Delete(item float64) error {
	for i, v := range *s {
		if v == item {
			(*s)[i] = (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
			return nil
		}
	}
	return errors.New("Item is not exists in Set")
}

// Test is Set implememntation test
func Test(s Set) {
	fmt.Printf("Adding 10 elements to set\n")
	for i := 0.0; i < 10; i++ {
		s.Add(i)
	}

	fmt.Printf("Is 5 contains in Set = %v\n", s.Contains(5.0))
	fmt.Printf("Is 10 contains in Set = %v\n", s.Contains(10.0))

	fmt.Printf("Deleting element 3\n")
	error := s.Delete(3.0)
	fmt.Printf("Error = %v\n", error)

	fmt.Printf("Deleting element 12\n")
	error = s.Delete(12.0)
	fmt.Printf("Error = %v\n\n\n", error)
}
