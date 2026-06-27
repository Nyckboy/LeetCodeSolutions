
// Time Complexity: O(N)
// Space Complexity: O(1)

/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function (s) {
  s = s.toLowerCase();
  let regularExp = /[a-zA-Z0-9]/;
  let left = 0;
  let right = s.length - 1;
  while (left < right) {
    if (regularExp.test(s[left]) && regularExp.test(s[right])) {
      if (s[left] != s[right]) return false;
      left++;
      right--;
    }else if (!regularExp.test(s[left])) {
      left++;
    }else if (!regularExp.test(s[right])) {
      right--;
    }
  }
  return true
};