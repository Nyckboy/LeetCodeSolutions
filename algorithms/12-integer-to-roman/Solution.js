
// Time Complexity: O(1)
// Space Complexity: O(1)

const romanDict = [
  ["M", 1000],
  ["CM", 900],
  ["D", 500],
  ["CD", 400],
  ["C", 100],
  ["XC", 90],
  ["L", 50],
  ["XL", 40],
  ["X", 10],
  ["IX", 9],
  ["V", 5],
  ["IV", 4],
  ["I", 1],
];

/**
 * @param {number} num
 * @return {string}
 */
// var intToRoman = function (num) {
//   let romNum = "";
//   while (num > 0) {
//     for (let i = 0; i < romanDict.length; i++) {
//       if (num >= romanDict[i][1]) {
//         romNum += romanDict[i][0];
//         num -= romanDict[i][1];
//         break;
//       }
//     }
//   }
//   return romNum
// };
var intToRoman = function (num) {
  let romNum = "";
  for (let i = 0; i < romanDict.length; i++) {
    while (num >= romanDict[i][1]) {
      romNum += romanDict[i][0];
      num -= romanDict[i][1];
    }
    if (num === 0) break;
  }
  return romNum
};

console.log(intToRoman(3749));
