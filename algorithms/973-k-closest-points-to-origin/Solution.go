
// Time Complexity: O(n log n) 
// Space Complexity: O(1)

package main

import (
	"fmt"
	"slices"
)

func kClosest(points [][]int, k int) [][]int {
	// Drop your logic here.
	// Target O(n log n) time, O(1) space.
	slices.SortFunc(points, func(a, b []int) int {
		return (a[0]*a[0] + a[1]*a[1]) - (b[0]*b[0] + b[1]*b[1])
	})
	
	return points[:k]
}

func main() {
	// Scenario A: Standard calculation
	// Expected: [[-2, 2]]
	// Reason: (-2)^2 + 2^2 = 8. (1)^2 + (3)^2 = 10. 8 is closer than 10.
	fmt.Println("Scenario A ->", kClosest([][]int{{1, 3}, {-2, 2}}, 1))

	// Scenario B: Multiple points
	// Expected: [[3, 3], [-2, 4]] (Order does not matter)
	fmt.Println("Scenario B ->", kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}