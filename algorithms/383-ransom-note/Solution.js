
// Time Complexity: O(n + m)
// Space Complexity: O(1)

/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
var canConstruct = function (ransomNote, magazine) {
  // Drop your logic here.
  // Target O(n + m) time, O(1) space.
  if (ransomNote.length > magazine.length) return false;
  let inv = {};
  for (const element of magazine) {
    if (inv[element]) {
      inv[element] += 1;
    } else {
      inv[element] = 1;
    }
  }
  for (const element of ransomNote) {
    if (inv[element]) {
      inv[element] -= 1;
    } else {
      return false;
    }
  }
  return true;
};

// --- Test Execution ---

// Scenario A: Standard Fail
// Note needs 'a', Magazine has 'b'
console.log("Scenario A (Expected: false) ->", canConstruct("a", "b"));

// Scenario B: Insufficient Quantity
// Note needs two 'a's, Magazine only has one
console.log("Scenario B (Expected: false) ->", canConstruct("aa", "ab"));

// Scenario C: Exact Match
// Note needs two 'a's, Magazine has two 'a's and a 'b'
console.log("Scenario C (Expected: true) ->", canConstruct("aa", "aab"));

// Scenario D: The Short-Circuit Edge Case
// Note is massive, Magazine is tiny
console.log("Scenario D (Expected: false) ->", canConstruct("abcdef", "abc"));
