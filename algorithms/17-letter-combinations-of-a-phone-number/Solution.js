
// Time Complexity: 
// Space Complexity: 

/**
 * @param {string} digits
 * @return {string[]}
 */

var letterCombinations = function(digits) {
    let array = [[], [], ["a", "b", "c"], ["d", "e", "f"], ["g", "h", "i"], ["j", "k", "l"], ["m", "n", "o"], ["p", "q", "r", "s"], ["t", "u", "v"], ["w", "x", "y", "z"]]
    let result = []
    digits = digits.split('');
    array[digits[0]*1].forEach(e1 => {
        if (digits.length >= 2) {
            array[digits[1]*1].forEach(e2 => {
                if (digits.length >= 3) {
                    array[digits[2]*1].forEach(e3 => {
                        if (digits.length >= 4) {
                            array[digits[3]*1].forEach(e4 => {
                                result.push(e1+e2+e3+e4)
                            })
                        }else{
                            result.push(e1+e2+e3)
                        }
                    })
                }else{
                    result.push(e1+e2)
                }
            })
        }else{
            result.push(e1)
        }
    });
    return result
};
