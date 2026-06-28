
// Time Complexity: O(N)
// Space Complexity: O(1)

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  if (s.length !== t.length) return false;
  let inv = {};
  for (let element of s) {
    if (inv[element]) {
      inv[element] += 1;
    } else {
      inv[element] = 1;
    }
  }
  for (let element of t) {
    if (inv[element]) {
      inv[element] -= 1;
    } else {
      return false;
    }
  }
  return true;
};

