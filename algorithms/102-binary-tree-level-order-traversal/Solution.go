
// Time Complexity: O(n)
// Space Complexity: O(n)

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// Drop your logic here.
	// Target O(n) time, O(n) space using a Queue Snapshot BFS.
	if root == nil { return [][]int{} }
	queue := []*TreeNode{root}
	var result [][]int
	
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}

func main() {
	// Scenario A: Standard Tree
	//      3
	//     / \
	//    9  20
	//      /  \
	//     15   7
	// Expected: [[3], [9, 20], [15, 7]]
	rootA := &TreeNode{Val: 3}
	rootA.Left = &TreeNode{Val: 9}
	rootA.Right = &TreeNode{Val: 20}
	rootA.Right.Left = &TreeNode{Val: 15}
	rootA.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Scenario A ->", levelOrder(rootA))

	// Scenario B: Single Node
	// Expected: [[1]]
	rootB := &TreeNode{Val: 1}
	fmt.Println("Scenario B ->", levelOrder(rootB))

	// Scenario C: Empty Tree
	// Expected: []
	var rootC *TreeNode = nil
	fmt.Println("Scenario C ->", levelOrder(rootC))
}