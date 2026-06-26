
// Time Complexity: O(N)
// Space Complexity: O(1)

/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  let buy = prices[0];
  let max = 0;
  for (let i = 0; i < prices.length; i++) {
    buy = Math.min(buy, prices[i]);
    max = Math.max(max, prices[i] - buy);
  }
  return max;
};

