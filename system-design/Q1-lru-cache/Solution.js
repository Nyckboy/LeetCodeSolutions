class Node {
  constructor(key, value) {
    this.key = key;
    this.value = value;
    this.prev = null; 
    this.next = null; 
  }
}

/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
    this.capacity = capacity;
    this.cache = new Map()
    this.head = new Node()
    this.tail = new Node()
    this.head.next = this.tail
    this.tail.prev = this.head
};

/** 
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
    if (!this.cache.has(key)) return -1
    let node = this.cache.get(key)
    this.moveToHead(node)
    return node.value
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
    if (!this.cache.has(key)) {
        let newNode = new Node(key,value)
        this.cache.set(key, newNode)
        newNode.next = this.head.next
        newNode.prev = this.head
        this.head.next.prev = newNode
        this.head.next = newNode
        if (this.cache.size > this.capacity) {
            let lastNode =  this.tail.prev
            this.cache.delete(lastNode.key)
            lastNode.prev.next = this.tail
            this.tail.prev = lastNode.prev
        }
    }else{
        let node = this.cache.get(key)
        node.value = value
        this.moveToHead(node)
    }
};

LRUCache.prototype.moveToHead = function(node) {
    node.prev.next = node.next
    node.next.prev = node.prev
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node 
}

/** 
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */
var obj = new LRUCache(10)
obj.get(10)
