package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	res := lexicalOrder2(2000)
	fmt.Println(res)
}

func lexicalOrder1(n int) []int {
	str := make([]string, 0, n)
	res := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		str = append(str, fmt.Sprint(i))
	}
	sort.Strings(str)
	for _, char := range str {
		num, _ := strconv.Atoi(char)
		res = append(res, num)
	}
	return res
}

func lexicalOrder2(n int) []int {
	result := make([]int, 0, n)
	for i := 1; i <= 9; i++ {
		dfs(i, n, &result)
	}
	return result
}

func dfs(current, n int, result *[]int) {
	if current > n {
		return
	}
	*result = append(*result, current)
	for i := 0; i <= 9; i++ {
		next := current*10 + i
		if next > n {
			return
		}
		dfs(next, n, result)
	}
}
