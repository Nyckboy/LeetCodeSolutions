
// Time Complexity: O(N * M) where N is haystack length and M is needle length
// Space Complexity: O(1)

/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function (haystack, needle) {
  // return haystack.indexOf(needle);
  let j = 0;
  for (let i = 0; i < haystack.length - needle.length; i++) {
    if (needle[j] != haystack[i]) continue;
    for (j = 0; j < needle.length; j++) {
        if (needle[j] == haystack[i + j]) continue;
        j = 0;
        break;
    }
    if (j != 0) return i;
  }
  return -1
};

console.log(strStr("sadbutsad", "sad"));
