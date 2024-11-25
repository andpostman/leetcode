package main

import "fmt"

func main() {
	fmt.Println(xorQueries1([]int{1, 3, 4, 8},
		[][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	fmt.Println(xorQueries2([]int{4, 8, 2, 10},
		[][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
}

func xorQueries1(arr []int, queries [][]int) []int {
	var resultArr []int
	for _, query := range queries {
		left, right := query[0], query[1]
		xorArr := arr[left : right+1]
		res := xorArr[0]
		for i := 1; i < len(xorArr); i++ {
			res ^= xorArr[i]
		}
		resultArr = append(resultArr, res)
	}
	return resultArr
}

func xorQueries2(arr []int, queries [][]int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] ^ arr[i]
	}

	ans := make([]int, 0)
	for _, query := range queries {
		x := arr[query[1]]
		if query[0] > 0 {
			x ^= arr[query[0]-1]
		}
		ans = append(ans, x)
	}

	return ans
}
