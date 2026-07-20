
// Time Complexity:  O(m * n)
// Space Complexity:  O(m * n)

package main

import (
	"fmt"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	// Drop your logic here.
	// Target O(m * n) time, O(m * n) space using Multi-Source BFS.
	rows := len(mat)
	cols := len(mat[0])
	queue := make([][2]int, 0, rows*cols)
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			} else {
				mat[r][c] = math.MaxInt32
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		r := curr[0]
		c := curr[1]

		for _, d := range dirs {
			nextR := r + d[0]
			nextC := c + d[1]
			if nextR >= 0 && nextR < rows && nextC >= 0 && nextC < cols  {
				if mat[nextR][nextC] > mat[r][c] + 1 {
					mat[nextR][nextC] = mat[r][c] + 1
					queue = append(queue, [2]int{nextR, nextC})
				}
			}
		}
	}
	return mat
}

func main() {
	// Scenario A: Standard Grid
	// Expected:
	// [[0, 0, 0],
	//  [0, 1, 0],
	//  [0, 0, 0]]
	matA := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println("Scenario A ->", updateMatrix(matA))

	// Scenario B: Deep Ripple
	// Expected:
	// [[0, 0, 0],
	//  [0, 1, 0],
	//  [1, 2, 1]]
	matB := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println("Scenario B ->", updateMatrix(matB))
}