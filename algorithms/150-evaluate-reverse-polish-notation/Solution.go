
// Time Complexity:  O(n) 
// Space Complexity:  O(n) 

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	// Drop your logic here.
	// Target O(n) time, O(n) space using a Stack.
	var stack []int
	for _, token := range tokens{
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)	
		} else {
			right := stack[len(stack) -1]
			stack = stack[:len(stack) -1]
			left := stack[len(stack) -1]
			stack = stack[:len(stack) -1]
			var result int
			switch token {
				case "+":
						result = left + right
				case "-":
						result = left - right
				case "*":
						result = left * right
				case "/":
						result = left / right
				}
			stack = append(stack, result)
		}
	}	
	return stack[0]
}

func main() {
	// Scenario A: Simple addition and multiplication
	// Expected: 9 (Explanation: (2 + 1) * 3)
	fmt.Println("Scenario A ->", evalRPN([]string{"2", "1", "+", "3", "*"}))

	// Scenario B: Division and truncation
	// Expected: 6 (Explanation: 4 + (13 / 5) -> 4 + 2)
	fmt.Println("Scenario B ->", evalRPN([]string{"4", "13", "5", "/", "+"}))

	// Scenario C: The Order Trap
	// Expected: 22
	// Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
	fmt.Println("Scenario C ->", evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}