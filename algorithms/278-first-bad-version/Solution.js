
// Time Complexity: O(log n)
// Space Complexity: O(1)

// --- The Mock API Factory ---
// This creates a fake isBadVersion function for testing
const createBadVersionAPI = (firstBad) => {
  return function (version) {
    return version >= firstBad;
  };
};

/**
 * Definition for isBadVersion()
 * * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function (isBadVersion) {
  /**
   * @param {integer} n Total versions
   * @return {integer} The first bad version
   */
  return function (n) {
    // Drop your logic here.
    // Target O(log n) time, O(1) space.
    let left = 1;
    let right = n;
    while (left < right) {
      let mid = left + Math.floor((right - left) / 2);
      if (isBadVersion(mid)) {
        right = mid;
      } else {
        left = mid + 1;
      }
    }
    return left;
  };
};

// --- Test Execution ---

// Scenario A: 5 versions, version 4 is the first bad one.
const apiA = createBadVersionAPI(4);
const findFirstBadVersionA = solution(apiA);
console.log("Scenario A (Expected: 4) ->", findFirstBadVersionA(5));

// Scenario B: 1 version, version 1 is bad.
const apiB = createBadVersionAPI(1);
const findFirstBadVersionB = solution(apiB);
console.log("Scenario B (Expected: 1) ->", findFirstBadVersionB(1));

// Scenario C: Massive scale, testing efficiency (no call stack overflow).
const apiC = createBadVersionAPI(70000000);
const findFirstBadVersionC = solution(apiC);
console.log(
  "Scenario C (Expected: 70000000) ->",
  findFirstBadVersionC(100000000),
);
