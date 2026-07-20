
// Time Complexity: O(n)
// Space Complexity: O(n)

package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	// Drop your logic here.
	// Target O(n) time, O(n) space.
	// Build it using the 3-phase sweep line architecture.
	n := len(intervals)
	i := 0
	result := make([][]int, 0, n+1)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}
	return result
}

func main() {
	// Scenario A: Clean insertion (No overlaps)
	// Expected: [[1, 2], [3, 4], [5, 6]]
	fmt.Println("Scenario A ->", insert([][]int{{1, 2}, {5, 6}}, []int{3, 4}))

	// Scenario B: The Collision Zone (Multiple overlaps)
	// Expected: [[1, 2], [3, 10], [12, 16]]
	fmt.Println("Scenario B ->", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))

	// Scenario C: Insert at the very beginning (Overlap with first)
	// Expected: [[1, 7]]
	fmt.Println("Scenario C ->", insert([][]int{{2, 5}, {6, 7}}, []int{1, 4}))

	// Scenario D: Envelops everything
	// Expected: [[1, 10]]
	fmt.Println("Scenario D ->", insert([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}}, []int{1, 10}))
}