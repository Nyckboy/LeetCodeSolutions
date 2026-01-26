
// Time Complexity: 
// Space Complexity: 

/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let isNegative = false
    if (x < 0) {
        x = x * -1
        isNegative = true 
    }
    let str = String(x).split('').reverse().join('')
    let num
    isNegative ? num = str*-1 : num = str*1;
    if (num < Math.pow(-2,31) || num > Math.pow(2,31) -1) {
        return 0
    }
    return num    
};

console.log(reverse(321));
