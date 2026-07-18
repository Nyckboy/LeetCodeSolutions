
// Time Complexity: O(n)
// Space Complexity: O(n)

package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	// Drop your logic here.
	// Target O(n) time, O(n) space using a zero-byte map.
	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}

func main() {
	// Scenario A: Standard Duplicate
	fmt.Println("Scenario A (Expected: true) ->", containsDuplicate([]int{1, 2, 3, 1}))

	// Scenario B: All Distinct
	fmt.Println("Scenario B (Expected: false) ->", containsDuplicate([]int{1, 2, 3, 4}))

	// Scenario C: Clustered Duplicates
	fmt.Println("Scenario C (Expected: true) ->", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}