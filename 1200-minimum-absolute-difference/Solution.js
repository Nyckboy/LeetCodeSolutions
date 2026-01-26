
// Time Complexity: 
// Space Complexity: 

/**
 * @param {number[]} arr
 * @return {number[][]}
 */
var minimumAbsDifference = function(arr) {
    arr.sort((a, b) => a - b)
    let result = []
    let min = 10000000000
    for (let i = 0; i < arr.length; i++) {
        if (arr[i+1] - arr[i] < min) {
            min = arr[i+1] - arr[i] 
        }
    }
    for (let i = 0; i < arr.length; i++) {
        if (arr[i+1] - arr[i] == min) {
            result.push([arr[i], arr[i+1]]) 
        }
    }
    return result
};

console.log(minimumAbsDifference([3,8,-10,23,19,-4,-14,27]));
