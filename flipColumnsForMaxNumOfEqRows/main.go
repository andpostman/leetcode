package main

import (
	"strings"
)

func main() {
	println(maxEqualRowsAfterFlips([][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}))
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	flipGrid := make(map[string]int)
	for _, col := range matrix {
		builder := strings.Builder{}
		for i := range col {
			if col[0] == col[i] {
				builder.WriteString("T")
			} else {
				builder.WriteString("F")
			}
		}
		flipGrid[builder.String()]++
	}
	maxNum := -1
	for _, val := range flipGrid {
		maxNum = max(val, maxNum)
	}
	return maxNum
}
