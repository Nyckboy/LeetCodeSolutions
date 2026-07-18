
// Time Complexity: O(n)
// Space Complexity: O(h)

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// Drop your logic here.
	// Target O(n) time, O(h) space (where h is the tree height).
	if root == nil { return 0	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func main() {
	// Scenario A: Standard Tree (Depth 3)
	//       3
	//      / \
	//     9  20
	//       /  \
	//      15   7
	rootA := &TreeNode{Val: 3}
	rootA.Left = &TreeNode{Val: 9}
	rootA.Right = &TreeNode{Val: 20}
	rootA.Right.Left = &TreeNode{Val: 15}
	rootA.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Scenario A (Expected: 3) ->", maxDepth(rootA))

	// Scenario B: Empty Tree
	var rootB *TreeNode = nil
	fmt.Println("Scenario B (Expected: 0) ->", maxDepth(rootB))
}