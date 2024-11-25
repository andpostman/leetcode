package main

import "fmt"

func main() {
	fmt.Println(maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
}

func maxMoves(grid [][]int) int {
	m := len(grid)    // number of rows
	n := len(grid[0]) // number of columns

	// res will store the rightmost column we can reach
	res := 0

	// dp array stores the maximum number of moves possible to reach each cell
	// in the current column we're processing
	dp := make([]int, m)

	// Iterate through each column from left to right (starting from column 1)
	for j := 1; j < n; j++ {
		// leftTop keeps track of the dp value from the cell above-left
		leftTop := 0
		// found indicates if we can reach any cell in current column
		found := false

		// Iterate through each row in current column
		for i := 0; i < m; i++ {
			// cur will store the maximum moves to reach current cell
			cur := -1
			// Store dp[i] for next iteration's leftTop
			nxtLeftTop := dp[i]

			// Check move from top-left cell (if valid)
			if i-1 >= 0 && leftTop != -1 && grid[i][j] > grid[i-1][j-1] {
				cur = max(cur, leftTop+1)
			}

			// Check move from direct left cell (if valid)
			if dp[i] != -1 && grid[i][j] > grid[i][j-1] {
				cur = max(cur, dp[i]+1)
			}

			// Check move from bottom-left cell (if valid)
			if i+1 < m && dp[i+1] != -1 && grid[i][j] > grid[i+1][j-1] {
				cur = max(cur, dp[i+1]+1)
			}

			// Update dp array for current cell
			dp[i] = cur
			// Update found flag if we can reach current cell
			found = found || (dp[i] != -1)
			// Update leftTop for next row's iteration
			leftTop = nxtLeftTop
		}

		// If we can't reach any cell in current column, break
		if !found {
			break
		}
		// Update result to current column if we can reach it
		res = j
	}

	// Return the maximum number of moves possible
	return res
}

// OR
func maxMoves2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	maxMoves := 0

	for col := n - 2; col >= 0; col-- {
		for row := 0; row < m; row++ {
			if row > 0 && grid[row][col] < grid[row-1][col+1] {
				dp[row][col] = max(dp[row][col], dp[row-1][col+1]+1)
			}
			if grid[row][col] < grid[row][col+1] {
				dp[row][col] = max(dp[row][col], dp[row][col+1]+1)
			}
			if row < m-1 && grid[row][col] < grid[row+1][col+1] {
				dp[row][col] = max(dp[row][col], dp[row+1][col+1]+1)
			}
		}
	}

	for row := 0; row < m; row++ {
		maxMoves = max(maxMoves, dp[row][0])
	}
	return maxMoves
}
