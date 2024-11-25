package main

import (
	"fmt"
)

func main() {
	println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
	println(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
}

func slidingPuzzle(board [][]int) int {
	goal := [][]int{{1, 2, 3}, {4, 5, 0}}
	move := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make(map[string]int)
	goalKey := fmt.Sprint(goal)
	if goalKey == fmt.Sprint(board) {
		return 0
	}
	var bfs func([][]int, int)
	bfs = func(board [][]int, count int) {
		key := fmt.Sprint(board)
		if goalKey == key {
			return
		}
		for i := range board {
			for j := range board[i] {
				if board[i][j] == 0 {
					for r := 0; r < len(move); r++ {
						newI := i + move[r][0]
						newJ := j + move[r][1]
						if newI >= 0 && newI < 2 && newJ >= 0 && newJ < 3 {
							newBoard := cp(board)
							newBoard[i][j], newBoard[newI][newJ] = board[newI][newJ], board[i][j]
							key = fmt.Sprint(newBoard)
							if val, ok := visited[key]; ok {
								if count+1 < val {
									visited[key] = count + 1
									bfs(newBoard, count+1)
								}
							} else {
								visited[key] = count + 1
								bfs(newBoard, count+1)
							}
						}
					}
					return
				}
			}
		}
	}
	bfs(board, 0)
	if _, ok := visited[goalKey]; !ok {
		return -1
	}
	return visited[goalKey]
}

func bfs(board [][]int, goal [][]int, move [][]int, visited map[string]int, count int) int {
	goalKey := fmt.Sprint(goal)
	key := fmt.Sprint(board)
	if goalKey == key {
		return count
	}
Loop:
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				for r := 0; r < len(move); r++ {
					newI := i + move[r][0]
					newJ := j + move[r][1]
					if newI >= 0 && newI < 2 && newJ >= 0 && newJ < 3 {
						newBoard := cp(board)
						newBoard[i][j], newBoard[newI][newJ] = board[newI][newJ], board[i][j]
						key = fmt.Sprint(newBoard)
						if val, ok := visited[key]; !ok || count+1 < val {
							visited[key] = count + 1
						}
						if key == goalKey {
							return visited[key]
						} else {
							bfs(newBoard, goal, move, visited, count+1)
						}

						// if val, ok := visited[key]; !ok || count+1 < val {
						// 	temp := bfs(newBoard, goal, move, visited, count+1)
						// 	if temp > 0 {
						// 		visited[key] = temp
						// 	}
						// }
					}
				}
				break Loop
			}
		}
	}
	return -1
}

func cp(board [][]int) [][]int {
	newBoard := make([][]int, len(board), cap(board))
	for l := 0; l < len(board); l++ {
		newBoard[l] = make([]int, len(board[l]), cap(board[l]))
		copy(newBoard[l], board[l])
	}
	return newBoard
}
