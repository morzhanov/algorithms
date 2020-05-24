package ds

import (
	"errors"
	"fmt"
)

// DoublyLinkedList implementation
// https://www.geeksforgeeks.org/data-structures/linked-list/
// Access O(n)
// Search O(n)
// Insert O(1)
// Delete O(1)
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

// DoublyLinkedListNode represents DoublyLinkedList node
type DoublyLinkedListNode struct {
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
	value float64
}

// Prepend value to array
func (d *DoublyLinkedList) Prepend(item float64) {
	node := DoublyLinkedListNode{value: item}
	if d.head == nil {
		d.tail = &node
	} else {
		node.next = d.head
		d.head.prev = &node
	}
	d.head = &node
}

// Append value to array
func (d *DoublyLinkedList) Append(item float64) {
	node := DoublyLinkedListNode{value: item}
	if d.head == nil {
		d.head = &node
	} else {
		d.tail.next = &node
	}
	d.tail = &node
	node.prev = d.tail
}

// InsertAfter method inserts a value after provided node
func (d *DoublyLinkedList) InsertAfter(item float64, node *DoublyLinkedListNode) error {
	if d.head == nil {
		return errors.New("List is empty")
	}

	newNode := DoublyLinkedListNode{value: item}
	newNode.next = node.next
	newNode.prev = node
	node.next.prev = &newNode
	node.next = &newNode
	return nil
}

// Delete node from list
func (d *DoublyLinkedList) Delete(node *DoublyLinkedListNode) error {
	if d.head == nil {
		return errors.New("List is empty")
	}
	if d.head == node {
		d.head = d.head.next
		d.head.prev = nil
		node.next = nil
		return nil
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if d.tail == node {
		d.tail = node.prev
	}
	node.next = nil
	node.prev = nil
	return nil
}

// Search value in list
func (d *DoublyLinkedList) Search(item float64) (*DoublyLinkedListNode, error) {
	if d.head == nil {
		return nil, errors.New("List is empty")
	}
	if d.head.value == item {
		return d.head, nil
	}
	if d.tail.value == item {
		return d.tail, nil
	}

	for node := d.head; node.next != nil; node = node.next {
		if node.value == item {
			return node, nil
		}
	}
	return nil, errors.New("Node not found")
}

// Update node value
func (d *DoublyLinkedList) Update(item float64, node *DoublyLinkedListNode) {
	node.value = item
}

// Traversal list
func (d *DoublyLinkedList) Traversal() []float64 {
	if d.head == nil {
		return nil
	}

	res := make([]float64, 0)
	for node := d.head; node.next != nil; node = node.next {
		res = append(res, node.value)
	}
	res = append(res, d.tail.value)
	return res
}

// TestDoublyLinkedList is DoublyLinkedList implememntation test
func TestDoublyLinkedList(l DoublyLinkedList) {
	fmt.Printf("DATA STRUCTURES: LINKED LIST \n")
	fmt.Printf("DLL Traversal = %v\n", l.Traversal())

	fmt.Printf("Adding 10 elements to list\n")
	for i := 0.0; i < 10; i++ {
		l.Append(i)
	}
	fmt.Printf("DLL Traversal = %v\n", l.Traversal())

	fmt.Printf("Search node with value 5\n")
	node, error := l.Search(5)
	fmt.Printf("Node next = %v, Error = %v\n", node.next, error)

	fmt.Printf("Inserting element after found node\n")
	error = l.InsertAfter(99, node)
	fmt.Printf("DLL Traversal = %v\n", l.Traversal())

	fmt.Printf("Updating node\n")
	l.Update(88, node)
	fmt.Printf("DLL Traversal = %v\n", l.Traversal())

	fmt.Printf("Deleting node\n")
	error = l.Delete(node)
	fmt.Printf("DLL Traversal = %v\n", l.Traversal())
}
