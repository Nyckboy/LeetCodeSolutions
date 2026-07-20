
// Time Complexity: O(n)
// Space Complexity: O(1)

package main

import "fmt"

func maxSubArray(nums []int) int {
	// Drop your logic here.
	// Target O(n) time, O(1) space using Kadane's Algorithm.
	currentSum := nums[0]
	maxSum := nums[0]
	for _, num := range nums[1:] {
		currentSum = max(currentSum + num, num)
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	// Scenario A: The standard mix of positive and negative
	// The subarray [4,-1,2,1] has the largest sum = 6.
	fmt.Println("Scenario A (Expected: 6) ->", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	// Scenario B: Single element
	fmt.Println("Scenario B (Expected: 1) ->", maxSubArray([]int{1}))

	// Scenario C: The Trap (All negative numbers)
	// The maximum subarray is just [-1] with sum -1.
	fmt.Println("Scenario C (Expected: -1) ->", maxSubArray([]int{-1, -2, -3, -4}))
}