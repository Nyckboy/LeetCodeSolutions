
// Time Complexity: 
// Space Complexity: 

/**
 * @param {number[]} height
 * @return {number}
 */

var maxArea = function(height) {
    let max = 0;
    let i = 0, j = height.length -1
    while (i < j) {
        let tmp = 0;
        height[i] < height[j] ? tmp = (j - i)*height[i] : tmp = (j - i)*height[j]
        if (max < tmp) {
            max = tmp
        } 
        height[i] < height[j] ? i++ : j--
    }
    return max
};

console.log(maxArea([1,1]));
