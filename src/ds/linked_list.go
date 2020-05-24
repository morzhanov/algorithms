package ds

import (
	"errors"
	"fmt"
	"strconv"
)

// LinkedList implementation
// https://www.geeksforgeeks.org/data-structures/linked-list/
// Access O(n)
// Search O(n)
// Insert O(1)
// Delete O(1)
type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

// LinkedListNode represents LinkedList node
type LinkedListNode struct {
	next  *LinkedListNode
	value LinkedListNodeItem
}

// LinkedListNodeItem is a LinkedListNode value
type LinkedListNodeItem struct {
	key   string
	value string
}

// Prepend value to array
func (d *LinkedList) Prepend(item LinkedListNodeItem) {
	node := LinkedListNode{value: item}
	if d.head == nil {
		d.tail = &node
	} else {
		node.next = d.head
	}
	d.head = &node
}

// Append value to array
func (d *LinkedList) Append(item LinkedListNodeItem) {
	node := LinkedListNode{value: item}
	if d.head == nil {
		d.head = &node
	} else {
		d.tail.next = &node
	}
	d.tail = &node
}

// InsertAfter method inserts a value after provided node
func (d *LinkedList) InsertAfter(item LinkedListNodeItem, node *LinkedListNode) error {
	if d.head == nil {
		return errors.New("List is empty")
	}

	newNode := LinkedListNode{value: item}
	newNode.next = node.next
	node.next = &newNode
	return nil
}

// Delete node from list
func (d *LinkedList) Delete(node *LinkedListNode) error {
	if d.head == nil {
		return errors.New("List is empty")
	}
	if d.head == node {
		d.head = d.head.next
		node.next = nil
		return nil
	}

	prev := d.head
	for prev.next != node {
		if prev.next == d.tail && d.tail != node {
			return errors.New("Node not found")
		}
		prev = prev.next
	}
	prev.next = node.next
	node.next = nil
	if d.tail == node {
		d.tail = prev
	}
	return nil
}

// Search value in list
func (d *LinkedList) Search(key string) (*LinkedListNode, error) {
	if d.head == nil {
		return nil, errors.New("List is empty")
	}
	if d.head.value.key == key {
		return d.head, nil
	}
	if d.tail.value.key == key {
		return d.tail, nil
	}

	for node := d.head; node.next != nil; node = node.next {
		if node.value.key == key {
			return node, nil
		}
	}
	return nil, errors.New("Node not found")
}

// Update node value
func (d *LinkedList) Update(item LinkedListNodeItem, node *LinkedListNode) {
	node.value = item
}

// Traversal list
func (d *LinkedList) Traversal() []LinkedListNodeItem {
	if d.head == nil {
		return nil
	}

	res := make([]LinkedListNodeItem, 0)
	for node := d.head; node.next != nil; node = node.next {
		res = append(res, node.value)
	}
	res = append(res, d.tail.value)
	return res
}

// Len method uses to get list length
func (d *LinkedList) Len() int {
	if d.head == nil {
		return 0
	}
	res := 0
	for node := d.head; node.next != nil; node = node.next {
		res++
	}
	if res != 0 {
		res++
	}
	return res
}

// TestLinkedList is LinkedList implememntation test
func TestLinkedList(l LinkedList) {
	fmt.Printf("DATA STRUCTURES: LINKED LIST \n")
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Adding 10 elements to list\n")
	for i := 0; i < 10; i++ {
		val := LinkedListNodeItem{key: strconv.Itoa(i), value: strconv.Itoa(i)}
		l.Append(val)
	}
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Search node with value 5\n")
	node, error := l.Search("5")
	fmt.Printf("Node next = %v, Error = %v\n", node.next, error)

	fmt.Printf("Inserting element after found node\n")
	val := LinkedListNodeItem{key: "99", value: "99"}
	error = l.InsertAfter(val, node)
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Updating node\n")
	val = LinkedListNodeItem{key: "55", value: "55"}
	l.Update(val, node)
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Deleting node\n")
	error = l.Delete(node)
	fmt.Printf("LL Traversal = %v\n", l.Traversal())
}
