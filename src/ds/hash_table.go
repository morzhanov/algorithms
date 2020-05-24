package ds

import (
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"
)

// HashTable implementation
// https://www.geeksforgeeks.org/java-util-hashmap-in-java-with-examples/
// Search: O(1)
// Insert: O(1)
// Delete: O(1)
// Space: O(n)
type HashTable struct {
	buckets map[HashTableKey]*LinkedList
}

// HashTableKey is a simple uint32 type
type HashTableKey uint32

// NewHashTable method used to create HashTable
func NewHashTable() *HashTable {
	var ht HashTable
	ht.buckets = make(map[HashTableKey]*LinkedList)
	return &ht
}

// Hash method uses to create hash from key
func (ht *HashTable) hash(key string) HashTableKey {
	h := fnv.New32a()
	h.Write([]byte(key))
	return HashTableKey(h.Sum32())
}

// Put method is used to put a value to HashTable
func (ht *HashTable) Put(key string, value string) {
	hash := ht.hash(key)
	bucket := ht.buckets[hash]

	val := LinkedListNodeItem{key: key, value: value}
	if bucket == nil {
		var ll LinkedList
		ll.Append(val)
		ht.buckets[hash] = &ll
	} else {
		node, error := ht.buckets[hash].Search(key)
		if error == nil {
			node.value = val
		} else {
			ht.buckets[hash].Append(val)
		}
	}
}

// Delete method is used to delete value from HashTable
func (ht *HashTable) Delete(key string) error {
	hash := ht.hash(key)
	bucket := ht.buckets[hash]

	if bucket == nil {
		return errors.New("Bucket not found")
	}

	node, error := ht.buckets[hash].Search(key)

	if error != nil {
		return errors.New("Hash key not found")
	}

	ht.buckets[hash].Delete(node)
	return nil
}

// Get method used to get value by Key from HashTable
func (ht *HashTable) Get(key string) (string, error) {
	hash := ht.hash(key)
	bucket := ht.buckets[hash]

	if bucket == nil {
		return "", errors.New("Bucket not found")
	}

	node, error := ht.buckets[hash].Search(key)

	if error != nil {
		return "", errors.New("Hash key not found")
	}
	return node.value.value, nil
}

// Has method used to check is key exists in HashTable
func (ht *HashTable) Has(key string) bool {
	hash := ht.hash(key)
	bucket := ht.buckets[hash]

	if bucket == nil {
		return false
	}

	_, error := ht.buckets[hash].Search(key)

	if error != nil {
		return false
	}
	return true
}

// Traversal returns key-value array of HashTable values
func (ht *HashTable) Traversal() []LinkedListNodeItem {
	var res []LinkedListNodeItem
	for _, v := range ht.buckets {
		res = append(res, v.Traversal()...)
	}
	return res
}

// TestHashTable is HashTable implememntation test
func TestHashTable(ht *HashTable) {
	fmt.Printf("DATA STRUCTURES: HASH TABLE \n")

	fmt.Printf("Adding 10 elements to table\n")
	for i := 0; i < 10; i++ {
		ht.Put(strconv.Itoa(i), strconv.Itoa(i*i))
	}
	fmt.Printf("HT Traversal = %v\n", ht.Traversal())

	fmt.Printf("Get value with key '5'\n")
	val, error := ht.Get("5")
	fmt.Printf("Val = %v, Error = %v\n", val, error)

	fmt.Printf("Deleting by key '5'\n")
	error = ht.Delete("5")
	fmt.Printf("HT Traversal = %v\n", ht.Traversal())
}
