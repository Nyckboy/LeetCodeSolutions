
// Time Complexity: O(N)
// Space Complexity: O(1)

package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	left := 1
	right := 1	
	for i := 0; i < n ; i++ {
		answer[i] = left
		left = left * nums[i] 
	}
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right = right * nums[i] 
	}
	return answer
}

func main() {
	// Scenario A: Standard array
	// Expected: [24 12 8 6]
	fmt.Println("Scenario A ->", productExceptSelf([]int{1, 2, 3, 4}))

	// Scenario B: Array with a zero
	// Expected: [0 0 9 0 0] 
	// (Everything becomes 0 except the spot where the 0 was!)
	fmt.Println("Scenario B ->", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}