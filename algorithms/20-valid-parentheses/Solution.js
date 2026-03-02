
// Time Complexity: 
// Space Complexity: 

/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    if(s.length % 2 !== 0) return false;
    var stack = []
    const map = {
      '{': '}',
      '[': ']',
      '(': ')',
    }
    for (let i = 0; i < s.length; i++) {
      const char = s[i];
      if(map[char]){
        stack.push(char);
      }else{
        let holder = stack.pop();
        if(map[holder] !== char) return false;
      }
    }
    return stack.length === 0;
};

