
// Time Complexity: O(N)
// Space Complexity: O(N)

"use strict";
/**
 * @param {number[][]} image
 * @param {number} sr
 * @param {number} sc
 * @param {number} color
 * @return {number[][]}
 */
var floodFill = function (image, sr, sc, color) {
  if (image[sr][sc] === color) return image;
  let originalColor = image[sr][sc];
  let rows = image.length;
  let cols = image[0].length;

  image[sr][sc] = color;
  if (sr + 1 < rows) {
    if (image[sr + 1][sc] === originalColor) {
      floodFill(image, sr + 1, sc, color);
    }
  }
  if (sr - 1 >= 0) {
    if (image[sr - 1][sc] === originalColor) {
      floodFill(image, sr - 1, sc, color);
    }
  }
  if (sc + 1 < cols) {
    if (image[sr][sc + 1] === originalColor) {
      floodFill(image, sr, sc + 1, color);
    }
  }
  if (sc - 1 >= 0) {
    if (image[sr][sc - 1] === originalColor) {
      floodFill(image, sr, sc - 1, color);
    }
  }
  return image;
};

console.log(
  floodFill(
    [
      [1, 1, 1],
      [1, 1, 0],
      [1, 0, 1],
    ],
    0,
    0,
    2,
  ),
);
