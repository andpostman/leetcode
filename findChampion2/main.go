package main

func main() {
	println(findChampion(3, [][]int{{0, 1}, {1, 2}}))
	println(findChampion(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
}

func findChampion(n int, edges [][]int) int {
	visited := make([]int, n)
	for _, edge := range edges {
		visited[edge[1]]++
	}
	ans := -1
	for i, num := range visited {
		if num == 0 && ans == -1 {
			ans = i
		} else if num == 0 {
			return -1
		}
	}
	return ans
}
