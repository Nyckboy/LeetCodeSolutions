
// Time Complexity: O(max(N, M))
// Space Complexity: O(max(N, M))

package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	// Drop your logic here.
	// Target O(max(N, M)) time, O(max(N, M)) space.

	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	result := make([]byte, maxLen + 1)

	i := len(a) - 1
	j := len(b) - 1
	k := maxLen
	carry := 0
	
	for i >= 0 || j >= 0 || carry > 0 {
    sum := carry 
		if i >= 0 {
        sum += int(a[i] - '0')
        i--                    
    }
    if j >= 0 {
        sum += int(b[j] - '0')
        j--
    }

		bit := sum%2
		carry = sum/2

		result[k] = byte(bit + '0')
		k--
	}
	if k == 0 {
		return string(result[1:])
	}
	
	return string(result)
}

func main() {
	// Scenario A: Standard addition
	fmt.Println("Scenario A (Expected: 100) ->", addBinary("11", "1"))

	// Scenario B: Equal length, cascading carry
	fmt.Println("Scenario B (Expected: 10101) ->", addBinary("1010", "1011"))

	// Scenario C: The Overflow Test (Would crash a 64-bit int)
	overflowA := "1111111111111111111111111111111111111111111"
	overflowB := "1"
	fmt.Println("Scenario C (Expected: 10000000000000000000000000000000000000000000) ->", addBinary(overflowA, overflowB))
}