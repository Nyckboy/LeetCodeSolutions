/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  let max = 0;
  let lastDup = 0;
  let subString = ""
    for (let i = 0; i < s.length; i++) {
        if(subString.includes(s[i])){
            max = subString.length > max ? subString.length : max;
            subString = "";
            i = lastDup++;
        }else{
            subString += s[i]
            max = subString.length > max ? subString.length : max;
        }
    }
    return max
};

console.log(lengthOfLongestSubstring("aab"));

/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    let longestStr = 0;
    let set = new Set()
    let left = 0;
    let right = 0;
    while(right < s.length){
        let char = s[right]
        if(!set.has(char)){
            set.add(char);
            longestStr = Math.max(longestStr,set.size);
            right++
        }else{
            set.delete(s[left]);
            left++
        }
    }
    return longestStr
};