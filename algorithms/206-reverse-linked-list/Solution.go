
// Time Complexity: O(N)
// Space Complexity: 0(1)

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 -> 2 -> 3 -> 4
// ^    ^
// 2 -> 1 -> 3 -> 4

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	// Drop your logic here.
	// Target O(n) time, O(1) space.
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// --- Helper function to print the list ---
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// Scenario A: 1 -> 2 -> 3 -> 4 -> 5 -> nil
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2}
	headA.Next.Next = &ListNode{Val: 3}
	headA.Next.Next.Next = &ListNode{Val: 4}
	headA.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Print("Original: ")
	printList(headA)
	
	reversedA := reverseList(headA)
	fmt.Print("Reversed: ")
	printList(reversedA) // Expected: 5 -> 4 -> 3 -> 2 -> 1 -> nil

	// Scenario B: Empty List (nil)
	var headB *ListNode = nil
	reversedB := reverseList(headB)
	fmt.Print("Reversed Empty: ")
	printList(reversedB) // Expected: nil
}