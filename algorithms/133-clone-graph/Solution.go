
// Time Complexity:  O(V + E)  (where V is vertices and E is edges)
// Space Complexity: O(V)

package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	// Drop your logic here.
	// Target O(V + E) time, O(V) space (where V is vertices and E is edges).
	if node == nil { return nil }
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node

	dfs = func(n *Node) *Node {
		if visited[n] != nil {
			return visited[n]
		} 
		clone := &Node{Val: n.Val}
		visited[n] = clone
		for _, neighbor := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor)) 
		}
		return clone
	}

	return dfs(node)
}

// --- Helper Functions for Testing ---
// Prints the graph as an adjacency list to verify structure
func printGraph(node *Node) {
	if node == nil {
		fmt.Println("[]")
		return
	}
	visited := make(map[int]bool)
	var printDFS func(*Node)
	printDFS = func(n *Node) {
		if visited[n.Val] {
			return
		}
		visited[n.Val] = true
		fmt.Printf("Node %d neighbors: [", n.Val)
		for i, neighbor := range n.Neighbors {
			fmt.Printf("%d", neighbor.Val)
			if i < len(n.Neighbors)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		for _, neighbor := range n.Neighbors {
			printDFS(neighbor)
		}
	}
	printDFS(node)
}

func main() {
	// Scenario A: A square graph (1-2-3-4 cycle)
	// 1 --- 2
	// |     |
	// 4 --- 3
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	fmt.Println("--- Original Graph ---")
	printGraph(node1)

	cloned := cloneGraph(node1)

	fmt.Println("\n--- Cloned Graph ---")
	printGraph(cloned)
    
	// Verification: The memory addresses should be different
	fmt.Printf("\nMemory Check: Original (%p) vs Clone (%p)\n", node1, cloned)
}