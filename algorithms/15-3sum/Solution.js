
// Time Complexity: $O(N^2)$
// Space Complexity: O(1)

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
  nums.sort((a, b) => a - b);
  let answer = [];
  for (let i = 0; i < nums.length - 2; i++) {
    let lock = nums[i];
    if (nums[i] == nums[i - 1] && i != 0) continue;
    let left = i + 1,
      right = nums.length - 1;
    while (left < right) {
      let sum = lock + nums[left] + nums[right];
      if (sum === 0) {
        answer.push([lock, nums[left], nums[right]]);
        left++;
        right--;
        while (left < right && nums[left] === nums[left - 1]) {
          left++;
        }
      } else if (sum > 0) {
        right--;
      } else {
        left++;
      }
    }
  }
  return answer;
};

