package dynamicarray

import (
	"errors"
	"fmt"
	"time"
)

// DynamicArray implementation
// https://www.geeksforgeeks.org/how-do-dynamic-arrays-work/
// Space O(n)
// Search O(1)
// Insert O(n)
// Delete O(n)
// Append O(1)
type DynamicArray struct {
	capacity int
	size     int
	items    []float64
}

func (d *DynamicArray) setSize(length int) {
	newItems := make([]float64, length)
	copy(newItems, d.items)

	d.items = newItems
	d.capacity = length
}

func (d *DynamicArray) resize(newSize int) {
	if newSize >= d.capacity {
		if d.capacity == 0 {
			d.capacity = 1
		}
		d.setSize(d.capacity * 2)
	} else if newSize < d.capacity/3 {
		d.setSize(d.capacity / 2)
	}
	d.size = newSize
}

// Append value to array
func (d *DynamicArray) Append(item float64) {
	d.resize(d.size + 1)
	d.items[d.size] = item
}

// Insert value to array
func (d *DynamicArray) Insert(item float64, position int) {
	d.resize(d.size + 1)

	for i := d.size - 1; i > position; i-- {
		d.items[i+1] = d.items[i]
	}
	d.items[position] = item
}

// Get value from array
func (d *DynamicArray) Get(position int) (float64, error) {
	if d.size < position {
		return 0, errors.New("Element is not exists")
	}
	return d.items[position], nil
}

// Delete value
func (d *DynamicArray) Delete(position int) error {
	if d.size < position {
		fmt.Printf("DELETE FAILED SIZE = %v POSITION = %v\n", d.size, position)
		return errors.New("Element is not exists")
	}

	d.resize(d.size - 1)

	for i := position; i < d.size-1; i++ {
		d.items[i] = d.items[i+1]
	}
	d.items[d.size-1] = 0
	return nil
}

// GetSize methods returns array size
func (d *DynamicArray) GetSize() int {
	return d.size
}

// GetCapacity methods returns array capacity
func (d *DynamicArray) GetCapacity() int {
	return d.capacity
}

// Test is Test implememntation test
func Test(da DynamicArray) {
	fmt.Printf("DA size: %v\n", da.GetSize())
	fmt.Printf("DA capacity: %v\n", da.GetCapacity())

	fmt.Printf("Adding 10 elements to array\n")
	for i := 0.0; i < 10; i++ {
		da.Append(i)
		fmt.Printf("DA size: %v\n", da.GetSize())
		fmt.Printf("DA capacity: %v\n", da.GetCapacity())
	}

	value, err := da.Get(5)
	fmt.Printf("Getting element from Array on index 5: %v. Error = %v\n", value, err)
	value, err = da.Get(11)
	fmt.Printf("Getting element from Array on index 11: %v. Error = %v\n", value, err)

	fmt.Printf("Deleting 10 elements from array\n")
	for i := 10; i > 1; i-- {
		da.Delete(i)
		fmt.Printf("DA size: %v\n", da.GetSize())
		fmt.Printf("DA capacity: %v\n", da.GetCapacity())
	}

	fmt.Printf("Adding 1K elements to array \n")
	before := time.Now()
	for i := 0.0; i < 1000; i++ {
		da.Append(i)
	}
	after := time.Now()
	fmt.Printf("DA size: %v\n", da.GetSize())
	fmt.Printf("DA capacity: %v\n", da.GetCapacity())
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Inserting element to array with 1K elements \n")
	before = time.Now()
	da.Insert(8.9, 500)
	after = time.Now()
	fmt.Printf("DA size: %v\n", da.GetSize())
	fmt.Printf("DA capacity: %v\n", da.GetCapacity())
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Deleting 1001 elements from array \n")
	before = time.Now()
	for i := 1001; i > 0; i-- {
		da.Delete(i)
		if i%100 == 0 {
			fmt.Printf("DA size: %v\n", da.GetSize())
			fmt.Printf("DA capacity: %v\n", da.GetCapacity())
		}
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Adding 1M elements to array \n")
	before = time.Now()
	for i := 0.0; i < 1000000; i++ {
		da.Append(i)
	}
	after = time.Now()
	fmt.Printf("DA size: %v\n", da.GetSize())
	fmt.Printf("DA capacity: %v\n", da.GetCapacity())
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Inserting element to array with 1M elements \n")
	before = time.Now()
	da.Insert(8.9, 500000)
	after = time.Now()
	fmt.Printf("DA size: %v\n", da.GetSize())
	fmt.Printf("DA capacity: %v\n", da.GetCapacity())
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Getting element from array with 1M elements \n")
	before = time.Now()
	value, err = da.Get(500000)
	after = time.Now()
	fmt.Printf("Val = %v\n", value)
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Deleting 1K + 1 elements from array \n")
	before = time.Now()
	for i := 1001; i > 0; i-- {
		da.Delete(i)
		if i%100 == 0 {
			fmt.Printf("DA size: %v\n", da.GetSize())
			fmt.Printf("DA capacity: %v\n", da.GetCapacity())
		}
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))
}
