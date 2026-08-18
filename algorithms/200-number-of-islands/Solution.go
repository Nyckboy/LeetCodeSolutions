
// Time Complexity: O(m*n)
// Space Complexity: O(m*n)

package main

import "fmt"

func numIslands(grid [][]byte) int {
	// Drop your logic here.
	// Target: O(m*n) time, O(m*n) space (call stack).
		
	// TODO: 1. Loop over every row 'r'
	// TODO: 2. Loop over every column 'c'
	// TODO: 3. If grid[r][c] == '1', increment count and call sinkIsland(grid, r, c)

	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				count++
				sinkIsland(grid, r, c)
			}
		}
	}	
	return count
}

func sinkIsland(grid [][]byte, r, c int) {
	row := len(grid)
	col := len(grid[0])

	// TODO: 1. Boundary check: return if out of bounds or if the cell is '0'
	// TODO: 2. "Sink" the current land by turning it into a '0'
	// TODO: 3. Recursively call sinkIsland for up, down, left, and right
	
	if r >= row || r < 0 || c >= col || c < 0 || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	sinkIsland(grid, r - 1 , c)
	sinkIsland(grid, r + 1, c)
	sinkIsland(grid, r, c - 1)
	sinkIsland(grid, r, c + 1)
}

func main() {
	// Scenario A: 1 large island
	gridA := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Scenario A ->", numIslands(gridA)) // Expected: 1

	// Scenario B: 3 separate islands
	gridB := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Scenario B ->", numIslands(gridB)) // Expected: 3
}