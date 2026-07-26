
// Time Complexity: O(n)
// Space Complexity: O(m * n)$ where m is the total number of words inserted, and n is their average length.

package main

import "fmt"

type Trie struct {
	// TODO: Define your children array and isEnd boolean here
	children [26]*Trie
	isEnd bool 
}

func Constructor() Trie {
	// TODO: Initialize and return a new Trie
	return Trie{}
}

func (this *Trie) Insert(word string) {
	// TODO: Traverse the word character by character.
	// If a child node doesn't exist for the character, create it.
	// Mark the very last node's isEnd as true.
	curr := this
	for _, char := range word {
		index := char - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}
		curr = curr.children[index]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	// TODO: Traverse the word. If you hit a nil child, return false.
	// If you reach the end of the word, return the node's isEnd value.
	curr := this
	for _, char := range word {
		index := char - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	// TODO: Exactly like Search, but if you successfully reach the 
	// end of the prefix without hitting a nil child, return true!
	curr := this
	for _, char := range prefix {
		index := char - 'a'
		if  curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return true
}

func main() {
	// Initialize the Trie
	trie := Constructor()

	// 1. Insert "apple"
	trie.Insert("apple")
	fmt.Println("Inserted 'apple'")

	// 2. Search exact word "apple"
	// Expected: true
	fmt.Println("Search 'apple' ->", trie.Search("apple"))

	// 3. Search exact word "app" (exists as prefix, but not as a standalone word yet)
	// Expected: false
	fmt.Println("Search 'app'   ->", trie.Search("app"))

	// 4. Check prefix "app"
	// Expected: true
	fmt.Println("Starts with 'app' ->", trie.StartsWith("app"))

	// 5. Insert "app" to make it an official word
	trie.Insert("app")
	fmt.Println("Inserted 'app'")

	// 6. Search exact word "app" again
	// Expected: true
	fmt.Println("Search 'app'   ->", trie.Search("app"))
}