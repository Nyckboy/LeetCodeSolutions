
// Time Complexity: O(log n)
// Space Complexity: O(1)

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// Kick off the recursion with no limits (-Infinity to +Infinity)
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min int, max int) bool {
	
	// TODO: 1. Base Case: If the node is nil, it's valid (return true).
	// TODO: 2. Boundary Check: If the node's value is <= min OR >= max, return false.
	// TODO: 3. Recursive Step: Check the Left branch AND the Right branch.
	// Hint for Left: The 'max' boundary becomes node.Val.
	// Hint for Right: The 'min' boundary becomes node.Val.
	// Return true ONLY if both branches return true.
	if node == nil { return true }
	if node.Val <= min || node.Val >= max { return false	}
	return validate(node.Left, min, node.Val) &&	validate(node.Right, node.Val, max)
}

func main() {
	// Scenario A: Valid BST
	//      2
	//     / \
	//    1   3
	// Expected: true
	rootA := &TreeNode{Val: 2}
	rootA.Left = &TreeNode{Val: 1}
	rootA.Right = &TreeNode{Val: 3}
	fmt.Println("Scenario A ->", isValidBST(rootA))

	// Scenario B: The Global Boundary Trap
	//      5
	//     / \
	//    1   4
	//       / \
	//      3   6
	// Expected: false (4 is smaller than 5, but 3 is in the right branch of 5!)
	rootB := &TreeNode{Val: 5}
	rootB.Left = &TreeNode{Val: 1}
	rootB.Right = &TreeNode{Val: 4}
	rootB.Right.Left = &TreeNode{Val: 3}
	rootB.Right.Right = &TreeNode{Val: 6}
	fmt.Println("Scenario B ->", isValidBST(rootB))
}