
// Time Complexity: 
// Space Complexity: 

class TrieNode {
    constructor() {
        this.children = {};
        this.isEnd = false;
    }
}

/**
 * @param {string[]} words
 */
var StreamChecker = function(words) {
    this.root = new TrieNode();
    this.stream = [];

    for (const word of words) {
        let node = this.root;
        for (let i = word.length -1; i >= 0; i--) {
            const char = word[i];
            if (!node.children[char]) {
                node.children[char] = new TrieNode();
            }
            node = node.children[char];
        }
        node.isEnd = true;
    }
};

/** 
 * @param {character} letter
 * @return {boolean}
 */
StreamChecker.prototype.query = function(letter) {
    this.stream.push(letter);
    let node = this.root;
    for (let i = this.stream.length -1; i >= 0; i--) {
        if (!node.children[this.stream[i]])break;
        node = node.children[this.stream[i]];
        if(node.isEnd === true) return true;
    }
    return false;
};

/** 
 * Your StreamChecker object will be instantiated and called as such:
 * var obj = new StreamChecker(words)
 * var param_1 = obj.query(letter)
 */