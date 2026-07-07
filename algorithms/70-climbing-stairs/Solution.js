
// Time Complexity: O(N)
// Space Complexity: O(1)

/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  // Drop your logic here.
  // Target O(n) time, O(1) space.
  if (n === 1) return 1;
  if (n === 2) return 2;

  let twoStepsBehind = 1;
  let oneStepBehind = 2;

  for (let i = 3; i <= n; i++) {
    let current = oneStepBehind + twoStepsBehind;
    twoStepsBehind = oneStepBehind;
    oneStepBehind = current;
  }
  return oneStepBehind;
};

// --- Test Execution ---

// Scenario A: The Base Case 1
// (1+1) or (2)
console.log("Scenario A (Expected: 2) ->", climbStairs(2));

// Scenario B: The Base Case 2
// (1+1+1), (1+2), or (2+1)
console.log("Scenario B (Expected: 3) ->", climbStairs(3));

// Scenario C: The Scale Test
// If your time complexity is O(2^n), this will take noticeably long or time out.
console.log("Scenario C (Expected: 8) ->", climbStairs(5));
