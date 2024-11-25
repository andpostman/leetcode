package main

import "fmt"

func main() {
	fmt.Println(rotateTheBox([][]byte{{'#', '.', '#'}}))
	fmt.Println(rotateTheBox([][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}))
}

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	for i := 0; i < m; i++ {
		link := make([]int, 0)
		for j := n - 1; j >= 0; j-- {
			switch box[i][j] {
			case '.':
				link = append(link, j)
			case '#':
				if len(link) == 0 {
					break
				}
				idx := link[0]
				if len(link) > 1 {
					link = link[1:]
				} else {
					link = []int{}
				}
				box[i][idx], box[i][j] = '#', '.'
				link = append(link, j)
			case '*':
				link = []int{}
			}
		}
	}
	rotatedBox := make([][]byte, n)
	for i := 0; i < n; i++ {
		rotatedBox[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			rotatedBox[i][j] = box[m-1-j][i]
		}
	}
	return rotatedBox
}

//second solution
func rotateTheBoxSec(box [][]byte) [][]byte {
	ROWS := len(box)
	COLS := len(box[0])

	res := make([][]byte, COLS)
	for i := range res {
		res[i] = make([]byte, ROWS)
		for j := range res[i] {
			res[i][j] = '.'
		}
	}

	for r := 0; r < ROWS; r++ {
		i := COLS - 1
		for c := COLS - 1; c >= 0; c-- {
			if box[r][c] == '#' {
				res[i][ROWS-r-1] = '#'
				i--
			} else if box[r][c] == '*' {
				res[c][ROWS-r-1] = '*'
				i = c - 1
			}
		}
	}
	return res
}
