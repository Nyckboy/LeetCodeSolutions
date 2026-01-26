
// Time Complexity: 
// Space Complexity: 

/**
 * @param {string[]} strs
 * @return {string} 
 */
var longestCommonPrefix = function(strs) {
    if( strs.length == 0 || strs[0] == "") return "";
    if ( strs.length == 1) return strs[0]
    let pref = "";
    let cpt = 0;
    let flage = 0;
    while (flage !=1) {
        if (strs[0][cpt] == undefined) break
        pref += strs[0][cpt++] || ""
        strs.forEach(e => {
            if (!e.startsWith(pref)) {
                flage = 1;
            }
        })
        if (flage == 1) {
            pref = pref.slice(0,-1)
        }
    }
    return pref
};