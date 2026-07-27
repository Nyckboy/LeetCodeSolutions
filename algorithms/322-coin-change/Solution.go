
// Time Complexity: O(amount * len(coins))
// Space Complexity: O(amount)

package main

import "fmt"

func coinChange(coins []int, amount int) int {
	// Drop your logic here.
	// Target O(amount * len(coins)) time, O(amount) space.
	dp := make([]int, amount + 1)
	maxVal := amount + 1
	for i := range dp {
		dp[i] = maxVal
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i - coin >= 0 {
				dp[i] = min(dp[i], dp[i - coin] + 1)
			}
		}
	}
	if dp[amount] == maxVal {
		return -1
	}
	return dp[amount]
}

func main() {
	// Scenario A: Standard case
	// Expected: 3 (11 = 5 + 5 + 1)
	fmt.Println("Scenario A ->", coinChange([]int{1, 2, 5}, 11))

	// Scenario B: The Greedy Trap
	// Expected: 2 (6 = 3 + 3)
	fmt.Println("Scenario B ->", coinChange([]int{1, 3, 4}, 6))

	// Scenario C: Impossible target
	// Expected: -1
	fmt.Println("Scenario C ->", coinChange([]int{2}, 3))

	// Scenario D: Zero amount
	// Expected: 0
	fmt.Println("Scenario D ->", coinChange([]int{1}, 0))
}