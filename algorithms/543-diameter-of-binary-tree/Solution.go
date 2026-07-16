
// Time Complexity: O(n)
// Space Complexity: O(h)

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	// Drop your logic here.
    // Target O(n) time, O(h) space (where h is the height of the tree for the call stack).
	maxDiam := 0;
	var dfs func(node *TreeNode) int;

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)

		if leftDepth + rightDepth > maxDiam {
			maxDiam = leftDepth + rightDepth
		}
		
		return 1 + max(leftDepth, rightDepth)
	}
	
	dfs(root)
	return maxDiam
}

// --- Helper for the Math.max equivalent ---
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Scenario A: Standard Tree
	//       1
	//      / \
	//     2   3
	//    / \     
	//   4   5    
	rootA := &TreeNode{Val: 1}
	rootA.Left = &TreeNode{Val: 2}
	rootA.Right = &TreeNode{Val: 3}
	rootA.Left.Left = &TreeNode{Val: 4}
	rootA.Left.Right = &TreeNode{Val: 5}
	
	fmt.Println("Scenario A (Expected: 3) ->", diameterOfBinaryTree(rootA))

	// Scenario B: Linear Tree
	//   1
	//  /
	// 2
	rootB := &TreeNode{Val: 1}
	rootB.Left = &TreeNode{Val: 2}

	fmt.Println("Scenario B (Expected: 1) ->", diameterOfBinaryTree(rootB))
}