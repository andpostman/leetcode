package main

func main() {
	println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 2
	}
	for _, guard := range guards {
		grid[guard[0]][guard[1]] = 2
	}
	for _, guard := range guards {
		guardX := guard[0]
		guardY := guard[1]
		for i := 0; i < len(directions); {
			newX := guardX + directions[i][0]
			newY := guardY + directions[i][1]
			if newX >= m || newY >= n || newX < 0 || newY < 0 || grid[newX][newY] == 2 {
				i++
				guardX = guard[0]
				guardY = guard[1]
			} else {
				grid[newX][newY] = 1
				guardX, guardY = newX, newY
			}
		}
	}
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				res++
			}
		}
	}
	return res
}
