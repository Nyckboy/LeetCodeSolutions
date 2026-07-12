package main

import "fmt"

func longestPalindrome(s string) int {
	// Drop your logic here.
	// Target O(n) time, O(1) space.
	var counts [128]int
	hasOdd := false
	length := 0

	for _, char := range s {
		counts[char]++
	}

	for _, count := range counts {
		length += (count / 2)*2
		if count % 2 == 1 && !hasOdd{
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}
	
	return length
}

func main() {
	// Scenario A: "abccccdd"
	// 4 'c's, 2 'd's, and one odd letter in the middle ('a' or 'b')
	fmt.Println("Scenario A (Expected: 7) ->", longestPalindrome("abccccdd"))

	// Scenario B: "a"
	fmt.Println("Scenario B (Expected: 1) ->", longestPalindrome("a"))

	// Scenario C: "bb"
	fmt.Println("Scenario C (Expected: 2) ->", longestPalindrome("bb"))
    
    // Scenario D: Case Sensitivity test "Aa"
	fmt.Println("Scenario D (Expected: 1) ->", longestPalindrome("Aa"))
}