
// Time Complexity: O(1)
// Space Complexity: O(N)

"use strict";
var MyQueue = function () {
  this.inbox = [];
  this.outbox = [];
};

/** * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function (x) {
  this.inbox.push(x);
};

/**
 * @return {number}
 */
MyQueue.prototype.pop = function () {
  this.peek();
  return this.outbox.pop();
};

/**
 * @return {number}
 */
MyQueue.prototype.peek = function () {
  if (this.outbox.length === 0 && this.inbox.length !== 0) {
    while (this.inbox.length > 0) {
      this.outbox.push(this.inbox.pop());
    }
  }
  let peek = this.outbox.pop();
  this.outbox.push(peek);
  return peek;
};

/**
 * @return {boolean}
 */
MyQueue.prototype.empty = function () {
  return this.outbox.length === 0 && this.inbox.length === 0;
};

// --- Test Execution ---
console.log("--- Initializing MyQueue ---");
var obj = new MyQueue();

console.log("Pushing 1...");
obj.push(1);
console.log("Pushing 2...");
obj.push(2);

console.log("Peek (Expected: 1) ->", obj.peek());
console.log("Pop (Expected: 1) ->", obj.pop());
console.log("Empty? (Expected: false) ->", obj.empty());

console.log("Pushing 3...");
obj.push(3);

console.log("Pop (Expected: 2) ->", obj.pop());
console.log("Pop (Expected: 3) ->", obj.pop());
console.log("Empty? (Expected: true) ->", obj.empty());