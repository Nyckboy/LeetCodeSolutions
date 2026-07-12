
// Time Complexity: O(N)
// Space Complexity: O(1)

package main

import "fmt"

func majorityElement(nums []int) int {
	// Drop your logic here.
	// Target O(n) time, O(1) space using Boyer-Moore Voting.
	var count int
	var candid int
	for _, num := range nums {
		if count == 0 {
			candid = num
		}
		if candid == num {
			count++
		}else {
			count--
		}	
	}
	return candid
}

func main() {
	// Scenario A: The majority is grouped at the start
	fmt.Println("Scenario A (Expected: 3) ->", majorityElement([]int{3, 3, 2}))

	// Scenario B: The majority is scattered
	fmt.Println("Scenario B (Expected: 2) ->", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

	// Scenario C: The single element edge case
	fmt.Println("Scenario C (Expected: 5) ->", majorityElement([]int{5}))
}