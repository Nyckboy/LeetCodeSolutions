
// Time Complexity: O(n^2)
// Space Complexity: O(n^2)

package main

import "fmt"

func predictTheWinner(nums []int) bool {
	n := len(nums)
	
	// 1. Create a 2D array dp of size n x n
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		// Base case: Subarrays of length 1. 
		// If i == j, the player just takes that single number.
		dp[i][i] = nums[i]
	}

	// 2. Loop over the lengths of the subarrays (from length 2 up to n)
	for length := 2; length <= n; length++ {
		// 3. Loop over the starting index 'i'
		for i := 0; i <= n-length; i++ {
			// Calculate the ending index 'j' based on the start and length
			j := i + length - 1
			
			// TODO: Calculate Choice Left
			// Take nums[i] and subtract the opponent's best future score for the rest (dp[i+1][j])
			left := nums[i] - dp[i+1][j]
			
			// TODO: Calculate Choice Right
			// Take nums[j] and subtract the opponent's best future score for the rest (dp[i][j-1])
			right := nums[j] - dp[i][j - 1]
			// TODO: Set dp[i][j] to the max of Choice Left and Choice Right
			// (You can use Go's built-in min/max functions!)
			dp[i][j] = max(left, right)
		}
	}

	// If the net score difference for the entire array (from 0 to n-1) is >= 0, Player 1 wins!
	return dp[0][n-1] >= 0
}

func main() {
	// Scenario A: The Greedy Trap
	// Expected: true (Player 1 takes 1, forcing Player 2 to expose 233)
	fmt.Println("Scenario A ->", predictTheWinner([]int{1, 5, 233, 7}))

	// Scenario B: P1 cannot win
	// Expected: false (P1 takes 1 -> P2 takes 5. P1 takes 2 -> P2 takes 7)
	fmt.Println("Scenario B ->", predictTheWinner([]int{1, 5, 2}))
}