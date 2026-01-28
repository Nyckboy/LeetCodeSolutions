
// Time Complexity: 
// Space Complexity: 

/**
 * @param {number} k
 * @param {number[]} nums
 */
var KthLargest = function(k, nums) {
    this.nums = nums.sort((a, b) =>  a - b)
    if(k < nums.length) this.nums = this.nums.slice(nums.length - k, nums.length) 
    this.k = k
};

/** 
 * @param {number} val
 * @return {number}
 */
KthLargest.prototype.add = function(val) {
    if(val < this.nums[0] && this.nums.length === this.k) return this.nums[0]
    this.nums.push(val)
    this.nums.sort((a, b) =>  a - b)
    if(this.nums.length > this.k) this.nums.shift()
    return this.nums[0]
};

/** 
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */