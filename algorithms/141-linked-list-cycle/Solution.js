"use strict";

// Time Complexity: O(N)
// Space Complexity: O(1)


/**
 * Definition for singly-linked list.
 */
function ListNode(val) {
  this.val = val;
  this.next = null;
}

// --- Scenario A: The Cycle ---
// 3 -> 2 -> 0 -> -4
//      ^          |
//      |__________|
const headA = new ListNode(3);
const nodeA2 = new ListNode(2);
const nodeA3 = new ListNode(0);
const nodeA4 = new ListNode(-4);

headA.next = nodeA2;
nodeA2.next = nodeA3;
nodeA3.next = nodeA4;
nodeA4.next = nodeA2; // The Cycle

// --- Scenario B: Clean Termination ---
// 1 -> 2 -> null
const headB = new ListNode(1);
headB.next = new ListNode(2);

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function (head) {
  // Drop your logic here.
  // Target O(n) time, O(1) space.
  let slow = head;
  let fast = head;

  while (fast !== null && fast.next !== null) {
    slow = slow.next;
    fast = fast.next.next;
    
    if (fast === slow) {
      return true;
    }
  }
  return false;
};

// --- Test Execution ---
console.log("Scenario A (Expected: true) ->", hasCycle(headA));
console.log("Scenario B (Expected: false) ->", hasCycle(headB));
console.log("Scenario C (Expected: false) ->", hasCycle(null));
