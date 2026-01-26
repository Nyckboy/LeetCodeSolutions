
// Time Complexity: 
// Space Complexity: 

/**
 * @param {string} s
 * @return {number}
 */
var myAtoi = function(s) {
    let str = s.trim();
    let isNegative = false
    let newStr = []
    for (let i = 0; i < str.length; i++) {
        if (i == 0 && (str[0] == '-' || str[0] =='+')) {
            if (str[0] == '-') {
                isNegative = true;          
            }
            continue;
        }
        if (str[i] >= "0" && str[i] <= "9") {
            newStr.push(str[i]);
        }else{
            break;
        }
    }
    let result = isNegative ? newStr.join('')*-1 : newStr.join('')*1;
    if (result > Math.pow(2, 31) -1) {
        return Math.pow(2, 31) -1;
    } else if(result < Math.pow(-2, 31)){
        return Math.pow(-2, 31);
    }
    return result;
};