# Data Structure: The Trie (Prefix Tree)

## 1. The Concept
A **Trie** (pronounced "try" or "tree") is a tree-based data structure used to store a collection of strings efficiently. Unlike an array where every word is stored separately, a Trie **merges common prefixes**.

**Think of it like a Folder System:**
* To store "Apple", you create nested folders: `A` -> `P` -> `P` -> `L` -> `E`.
* To store "Apply", you use the **same** `A` -> `P` -> `P` -> `L` path, but then branch into a new folder `Y`.



## 2. Why use it?
* **Space Efficiency:** You don't store "App" twice. You store the prefix once and branch off only where words differ.
* **Speed:** Checking if a word exists takes **O(L)** time (where $L$ is the length of the word you are looking for).
    * It does **not** matter if the dictionary has 10 words or 10 million words; the lookup speed depends only on the length of the word itself.

---

## 3. The Implementation (Standard Template)

Here is the standard boilerplate for a Trie in JavaScript.

```javascript
/**
 * 1. The Node
 * Each node represents a single character in the tree.
 */
class TrieNode {
    constructor() {
        this.children = {}; // Map of character -> TrieNode
        this.isEnd = false; // Flag: Does a valid word finish exactly here?
    }
}

/**
 * 2. The Trie Class
 */
class Trie {
    constructor() {
        this.root = new TrieNode();
    }

    /**
     * Inserts a word into the Trie.
     * Time: O(L) where L is word length.
     */
    insert(word) {
        let node = this.root;
        for (const char of word) {
            // If the folder/path doesn't exist, create it
            if (!node.children[char]) {
                node.children[char] = new TrieNode();
            }
            // Move into the child folder
            node = node.children[char];
        }
        // Mark the end of the word
        node.isEnd = true;
    }

    /**
     * Checks if a full word exists in the Trie.
     * Time: O(L)
     */
    search(word) {
        let node = this.root;
        for (const char of word) {
            // If the path breaks, the word doesn't exist
            if (!node.children[char]) return false;
            node = node.children[char];
        }
        // It's only a match if we marked this node as the end of a word
        return node.isEnd;
    }

    /**
     * Checks if there is ANY word in the Trie that starts with the given prefix.
     * Time: O(L)
     */
    startsWith(prefix) {
        let node = this.root;
        for (const char of prefix) {
            if (!node.children[char]) return false;
            node = node.children[char];
        }
        return true;
    }
}
```

## 4. Advanced Strategy: The "Reverse Trie"

**Use Case:** Problems involving **Suffixes** or **Stream History**.
* *Example Problem:* "Does the stream of characters I just typed end with a known word?" (e.g., Stream `...abcde`, Dictionary has `cde`).

**The Logic:**
Instead of storing words normally (`c -> d -> e`), store them **Backwards** (`e -> d -> c`).

1.  **Insert:** Reverse every dictionary word before inserting it into the Trie.
2.  **Query:** Walk **backwards** through your stream history (current char, previous char, etc.) while walking down the Trie.

**Why?**
* Comparing a suffix against a list of words is slow ($O(N \times Words)$).
* Walking a Reverse Trie is fast ($O(WordLength)$). You stop immediately when you find a match or hit a dead end.

---

## 5. Summary Checklist
* [ ] **Node Structure:** Does it have `children = {}` and `isEnd = false`?
* [ ] **Insert Logic:** Are you creating new nodes when `!node.children[char]`?
* [ ] **Search Logic:** Are you checking `node.isEnd` at the very end?
* [ ] **Prefix Logic:** For `startsWith`, you don't check `isEnd`—just if the path exists.
* [ ] **Suffix Problems:** If checking suffixes, build the Trie with **Reversed Words**.