
// Time Complexity: O(V + E)
// Space Complexity: O(V + E)
// V = The number of courses (numCourses).
// E = The number of prerequisites (the length of the prerequisites array).

package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. Build the Adjacency List (Directed Graph)
	adj := make([][]int, numCourses)
	for _, pre := range prerequisites {
		course := pre[0]
		prereq := pre[1]
		// An edge from prereq to course (prereq -> course)
		adj[prereq] = append(adj[prereq], course)
	}
	// 2. The 3-State Ledger
	// 0 = Unvisited, 1 = Visiting (Active Path), 2 = Safe (Fully Checked)
	state := make([]int, numCourses)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		// TODO: 1. If state[node] is 1 (Visiting), we found a cycle! Return false.
		if state[node] == 1 { return false }
		// TODO: 2. If state[node] is 2 (Safe), we are good. Return true.
		if state[node] == 2 { return true }
		
		// TODO: 3. Mark the current node as 1 (Visiting).
		state[node] = 1
		
		// TODO: 4. Loop through adj[node]. 
		//       For each neighbor, call dfs(neighbor). 
		//       If any neighbor returns false, immediately return false.
		for _, neighbor := range adj[node] {
			if !dfs(neighbor) { return false }
		}
		
		// TODO: 5. We finished checking all neighbors safely! 
		//       Mark state[node] as 2 (Safe) and return true.
			state[node] = 2
			return true
	}

	// 3. We must check EVERY node, because the graph might be disconnected
	//    (e.g., two completely separate degree tracks).
	for i := 0; i < numCourses; i++ {
		if state[i] == 0 {
			// If DFS returns false, a cycle exists.
			if !dfs(i) {
				return false
			}
		}
	}

	return true
}

func main() {
	// Scenario A: Safe progression (1 -> 0)
	// Expected: true
	fmt.Println("Scenario A ->", canFinish(2, [][]int{{1, 0}}))

	// Scenario B: Direct Cycle (1 -> 0, and 0 -> 1)
	// Expected: false
	fmt.Println("Scenario B ->", canFinish(2, [][]int{{1, 0}, {0, 1}}))

	// Scenario C: The Diamond (Safe, but tricky for a basic true/false map)
	// 0 requires 1 and 2. Both 1 and 2 require 3. (3 -> 1 -> 0, and 3 -> 2 -> 0)
	// Expected: true
	fmt.Println("Scenario C ->", canFinish(4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}))
}