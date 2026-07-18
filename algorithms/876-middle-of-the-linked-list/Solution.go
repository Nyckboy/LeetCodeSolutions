
// Time Complexity: O(n)
// Space Complexity: O(1)

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	// Drop your logic here.
	// Target O(n) time, O(1) space, single pass.
	// if head == nil { return nil }
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	
	return slow
}

func main() {
	// Scenario A: Odd Number of Nodes (1 -> 2 -> 3 -> 4 -> 5 -> nil)
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2}
	headA.Next.Next = &ListNode{Val: 3}
	headA.Next.Next.Next = &ListNode{Val: 4}
	headA.Next.Next.Next.Next = &ListNode{Val: 5}

	resA := middleNode(headA)
	fmt.Println("Scenario A (Expected: 3) ->", resA.Val)

	// Scenario B: Even Number of Nodes (1 -> 2 -> 3 -> 4 -> 5 -> 6 -> nil)
	// Should return the second middle node (4)
	headB := &ListNode{Val: 1}
	headB.Next = &ListNode{Val: 2}
	headB.Next.Next = &ListNode{Val: 3}
	headB.Next.Next.Next = &ListNode{Val: 4}
	headB.Next.Next.Next.Next = &ListNode{Val: 5}
	headB.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	resB := middleNode(headB)
	fmt.Println("Scenario B (Expected: 4) ->", resB.Val)
}