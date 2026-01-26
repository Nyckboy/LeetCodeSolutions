
// Time Complexity: 
// Space Complexity: 


//  Definition for singly-linked list.
 function ListNode(val, next) {
     this.val = (val===undefined ? 0 : val)
     this.next = (next===undefined ? null : next)
 }
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    let head = new ListNode(0)
    let tmp = head;
    let carry = 0
    while(l1 != null || l2 != null || carry != 0){
        let num1 = l1 == null ? 0 : l1.val;
        let num2 = l2 == null ? 0 : l2.val;
        let sum = num1 + num2 + carry;
        let digit = sum % 10
        carry = sum >= 10 ? 1 : 0;
        tmp.next = new ListNode(digit)
        tmp = tmp.next        
        l1 = (l1 != null) ? l1.next : null;
        l2 = (l2 != null) ? l2.next : null;
    }
    return head.next;
};

let l1 = new ListNode(2, new ListNode(4, new ListNode(3)))
let l2 = new ListNode(5, new ListNode(6, new ListNode(4)))


console.log(addTwoNumbers(l1, l2));
