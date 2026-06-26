
// Time Complexity: O(log N)
// Space Complexity: O(1)

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function (nums, target) {
  let left = 0,
    right = nums.length - 1;
  let middle 
  while (left <= right) {
    middle = parseInt((right + left ) /2 )    
    if (nums[middle] === target) return middle;
    if (nums[middle] < target) {
      left = middle + 1;
    } else {
      right = middle - 1;
    }
  }
  return left;
};



