package ds

import "fmt"

const alphabetSize = 26

// TrieNode is a Trie node implementation
type TrieNode struct {
	children    []*TrieNode
	isEndOfWord bool
}

// Trie implements Trie
type Trie struct {
	root *TrieNode
}

// CreateTrieNode function creates trie node
func CreateTrieNode() *TrieNode {
	var tn TrieNode
	tn.children = make([]*TrieNode, alphabetSize)
	return &tn
}

// CreateTrie function creates trie
func CreateTrie() *Trie {
	var t Trie
	t.root = CreateTrieNode()
	return &t
}

// IsEmpty function returns if Trie empty
func (t *Trie) IsEmpty() bool {
	for i := 0; i < alphabetSize; i++ {
		if t.root.children[i] != nil {
			return false
		}
	}
	return true
}

// IsEmpty function returns if Trie Node Tree is empty
func (n *TrieNode) IsEmpty() bool {
	for i := 0; i < alphabetSize; i++ {
		if n.children[i] != nil {
			return false
		}
	}
	return true
}

// Delete function removes Trie nodes from Trie by key provided
// returns new Trie root
func (t *Trie) Delete(root *TrieNode, key string, depth int) *TrieNode {
	chars := []rune(key)
	if root == nil {
		root = nil
		return nil
	}
	// if last character of key is being processed
	if depth == len(chars) {
		if root.isEndOfWord {
			root.isEndOfWord = false
		}
		if root.IsEmpty() {
			root = nil
			return nil
		}
		return root
	}

	idx := chars[depth] - 'a'
	root.children[idx] = t.Delete(root.children[idx], key, depth+1)

	if root.IsEmpty() && root.isEndOfWord == false {
		root = nil
		return nil
	}
	return root
}

// Insert function inserts a key into trie if not present
// If key is prefix of trie node, just ,arks leaf node
func (t *Trie) Insert(key string) {
	chars := []rune(key)
	length := len(key)
	crawl := t.root

	for level := 0; level < length; level++ {
		idx := chars[level] - 'a'
		if crawl.children[idx] == nil {
			crawl.children[idx] = CreateTrieNode()
		}
		crawl = crawl.children[idx]
	}
	crawl.isEndOfWord = true
}

// Search function returns true if key presents in trie, else false
func (t *Trie) Search(key string) bool {
	chars := []rune(key)
	length := len(key)
	crawl := t.root

	for level := 0; level < length; level++ {
		idx := chars[level] - 'a'
		if crawl.children[idx] == nil {
			return false
		}
		crawl = crawl.children[idx]
	}
	return crawl != nil && crawl.isEndOfWord
}

// TraverseTrieItem represents traverse element of Trie node
type TraverseTrieItem struct {
	key      string
	children []*TraverseTrieItem
}

// Traverse function traverses trie
func (t *Trie) Traverse() []*TraverseTrieItem {
	fmt.Printf("Traversing Trie\n")
	res := t.root.Traverse()
	for _, item := range res {
		printTraverseTrieItem(item)
	}
	fmt.Printf("\n")
	return res
}

func printTraverseTrieItem(item *TraverseTrieItem) {
	fmt.Printf("%v ", item.key)
	for _, child := range item.children {
		printTraverseTrieItem(child)
	}
}

// Traverse function traverses trie node tree
func (t *TrieNode) Traverse() []*TraverseTrieItem {
	nodes := make([]*TraverseTrieItem, 0)

	for idx, item := range t.children {
		if item != nil {
			node := TraverseTrieItem{key: string('a' + idx)}
			if item.children != nil {
				children := item.Traverse()
				if len(children) > 0 {
					node.children = children
				}
			}
			nodes = append(nodes, &node)
		}
	}
	return nodes
}

// TestTrie function tests Trie
func TestTrie(t *Trie) {
	keys := []string{"the", "a", "there", "answer", "any", "by", "bye", "their"}
	checkKeys := []string{"the", "these", "their", "thaw", "b", "a"}

	for i := 0; i < len(keys); i++ {
		t.Insert(keys[i])
	}
	t.Traverse()
	for i := 0; i < len(checkKeys); i++ {
		if t.Search(checkKeys[i]) {
			fmt.Printf("\"%v\" key is present in Trie\n", checkKeys[i])
		} else {
			fmt.Printf("\"%v\" key is not present in Trie\n", checkKeys[i])
		}
	}

	fmt.Printf("Removing key \"answer\" from Trie\n")
	t.root = t.Delete(t.root, "answer", 0)
	t.Traverse()
}
