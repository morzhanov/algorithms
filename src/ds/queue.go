package ds

import (
	"fmt"
	"time"
)

// Queue implementation
// https://www.geeksforgeeks.org/queue-data-structure/
// Access: O(n)
// Search: O(n)
// Insert: O(1)
// Delete: O(1)
type Queue []float64

// IsEmpty returns is Queue empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Len returns Queue length
func (q *Queue) Len() int {
	return len(*q)
}

// Enqueue is used to push value to Queue
func (q *Queue) Enqueue(item float64) {
	*q = append(*q, item)
}

// Dequeue is used to shift value from Queue, return true if queue is empty
func (q *Queue) Dequeue() (float64, bool) {
	if q.IsEmpty() {
		return 0, true
	}

	res := (*q)[0]
	*q = (*q)[1:]
	return res, false
}

// Front returns a first value from Queue, return true if queue is empty
func (q *Queue) Front() (float64, bool) {
	if q.IsEmpty() {
		return 0, true
	}
	return (*q)[0], false
}

// TestQueue is Queue implememntation test
func TestQueue(q Queue) {
	fmt.Printf("DATA STRUCTURES: QUEUE \n")
	fmt.Printf("Queue length: %v\n", q.Len())

	fmt.Printf("Pushing to queue %v\n", 10)
	q.Enqueue(10)
	fmt.Printf("Pushing to queue %v\n", 20)
	q.Enqueue(20)
	top, empty := q.Front()
	fmt.Printf("Front value = %v. Queue empty = %v\n", top, empty)
	fmt.Printf("Queue length = %v\n", q.Len())

	fmt.Printf("Dequeue value from queue\n")
	value, empty := q.Dequeue()
	fmt.Printf("Value = %v. Queue empty = %v\n", value, empty)
	fmt.Printf("Dequeue value from queue\n")
	value, empty = q.Dequeue()
	fmt.Printf("Value = %v. Queue empty = %v\n", value, empty)

	top, empty = q.Front()
	fmt.Printf("Front value = %v. Queue empty = %v\n", top, empty)
	fmt.Printf("Queue length = %v\n", q.Len())

	fmt.Printf("Adding 1K elements\n")
	before := time.Now()
	for i := 0.0; i < 1000; i++ {
		q.Enqueue(i)
	}
	after := time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Extracting all 1K elements\n")
	before = time.Now()
	for i := 0.0; i < 1000; i++ {
		q.Dequeue()
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Adding 100M elements\n")
	before = time.Now()
	for i := 0.0; i < 100000000; i++ {
		q.Enqueue(i)
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))

	fmt.Printf("Extracting all 100M elements\n")
	before = time.Now()
	for i := 0.0; i < 100000000; i++ {
		q.Dequeue()
	}
	after = time.Now()
	fmt.Printf("Time taken: %v\n\n\n", after.Sub(before))
}
