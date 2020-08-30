package linkedlist

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
	head *Node
	tail *Node
}

// Node represents LinkedList node
type Node struct {
	next  *Node
	value NodeItem
}

// SetValue sets Node value
func (n *Node) SetValue(val NodeItem) {
	n.value = val
}

// GetValue returns Node value
func (n *Node) GetValue() NodeItem {
	return n.value
}

// NodeItem is a LinkedList node value
type NodeItem struct {
	key   string
	value string
}

// CreateNodeItem creates node item
func CreateNodeItem(key string, value string) NodeItem {
	return NodeItem{key: key, value: value}
}

// GetKey returns NodeItem key
func (n *NodeItem) GetKey() string {
	return n.key
}

// GetValue returns NodeItem value
func (n *NodeItem) GetValue() string {
	return n.value
}

// SetValue sets NodeItem value
func (n *NodeItem) SetValue(val string) {
	n.value = val
}

// Prepend value to array
func (d *LinkedList) Prepend(item NodeItem) {
	node := Node{value: item}
	if d.head == nil {
		d.tail = &node
	} else {
		node.next = d.head
	}
	d.head = &node
}

// Append value to array
func (d *LinkedList) Append(item NodeItem) {
	node := Node{value: item}
	if d.head == nil {
		d.head = &node
	} else {
		d.tail.next = &node
	}
	d.tail = &node
}

// InsertAfter method inserts a value after provided node
func (d *LinkedList) InsertAfter(item NodeItem, node *Node) error {
	if d.head == nil {
		return errors.New("List is empty")
	}

	newNode := Node{value: item}
	newNode.next = node.next
	node.next = &newNode
	return nil
}

// Delete node from list
func (d *LinkedList) Delete(node *Node) error {
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
func (d *LinkedList) Search(key string) (*Node, error) {
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
func (d *LinkedList) Update(item NodeItem, node *Node) {
	node.value = item
}

// Traversal list
func (d *LinkedList) Traversal() []NodeItem {
	if d.head == nil {
		return nil
	}

	res := make([]NodeItem, 0)
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

// Test is LinkedList implememntation test
func Test(l LinkedList) {
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Adding 10 elements to list\n")
	for i := 0; i < 10; i++ {
		val := NodeItem{key: strconv.Itoa(i), value: strconv.Itoa(i)}
		l.Append(val)
	}
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Search node with value 5\n")
	node, error := l.Search("5")
	fmt.Printf("Node next = %v, Error = %v\n", node.next, error)

	fmt.Printf("Inserting element after found node\n")
	val := NodeItem{key: "99", value: "99"}
	error = l.InsertAfter(val, node)
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Updating node\n")
	val = NodeItem{key: "55", value: "55"}
	l.Update(val, node)
	fmt.Printf("LL Traversal = %v\n", l.Traversal())

	fmt.Printf("Deleting node\n")
	error = l.Delete(node)
	fmt.Printf("LL Traversal = %v\n", l.Traversal())
}
