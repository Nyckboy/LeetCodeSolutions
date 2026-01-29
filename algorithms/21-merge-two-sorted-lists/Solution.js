
// Time Complexity: 
// Space Complexity: 

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function(list1, list2) {
    if(list1 === null && list2 === null) return null;
    if(list1 === null && list2 !== null) return list2
    if(list1 !== null && list2 === null) return list1
    let list = new ListNode()
    tmp = list
    let flag = 0
    while (true) {
        const val1 = list1.val
        const val2 = list2.val
        if (val1 < val2) {
            tmp.val = val1
            list1 = list1.next? list1.next : flag = 1  
        }else{
            tmp.val = val2
            list2 = list2.next? list2.next : flag = 2  
        }
        if(flag != 0) break;
        tmp.next = new ListNode()
        tmp = tmp.next
    }
    if (flag == 1) tmp.next = list2
    if (flag == 2) tmp.next = list1
    return list
};