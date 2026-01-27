
// Time Complexity: O(1)
// Space Complexity: 

class Node {
  constructor(key, value) {
    this.key = key;
    this.value = value;
    this.freq = 1;
    this.prev = null; 
    this.next = null; 
  }
}

class MyDoublyLinkedList {
    constructor() {
        this.head = new Node();
        this.tail = new Node();
        this.head.next = this.tail;
        this.tail.prev = this.head;
        this.size = 0;
    }

    addNode(node) {
        node.next = this.head.next;
        node.prev = this.head;
        this.head.next.prev = node;
        this.head.next = node;
        this.size++;
    }

    removeNode(node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
        this.size--;
    }
}

/**
 * @param {number} capacity
 */
var LFUCache = function(capacity) {
    this.capacity = capacity;
    this.cache = new Map();
    this.freqMap = new Map();
    this.freqMap.set(1, new MyDoublyLinkedList());
    this.minFreq = 1;
};

/** 
 * @param {number} key
 * @return {number}
 */
LFUCache.prototype.get = function(key) {
    if (!this.cache.has(key)) return -1;
    let node = this.cache.get(key);
    this.updateFreq(node);
    return node.value;
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LFUCache.prototype.put = function(key, value) {
    if (this.capacity === 0) return;
    if (!this.cache.has(key)) {
        if (this.cache.size === this.capacity) {
            let list = this.freqMap.get(this.minFreq);
            this.cache.delete(list.tail.prev.key);
            list.removeNode(list.tail.prev);
        }
        let newNode = new Node(key,value);
        let list = this.freqMap.get(1);
        list.addNode(newNode);
        this.cache.set(key, newNode);
        this.minFreq = 1;
    }else{
        let node = this.cache.get(key);
        node.value = value;
        this.updateFreq(node)
    }
};

LFUCache.prototype.updateFreq = function(node) {
    let list2 = this.freqMap.get(this.minFreq);
    this.freqMap.get(node.freq).removeNode(node);
    if (node.freq === this.minFreq && list2.size === 0) {
        this.minFreq +=1;
    }
    node.freq += 1;
    if (!this.freqMap.has(node.freq)) {
        this.freqMap.set(node.freq, new MyDoublyLinkedList());
    }
    let list = this.freqMap.get(node.freq);
    list.addNode(node);
}

/** 
 * Your LFUCache object will be instantiated and called as such:
 * var obj = new LFUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */