package main

import "fmt"

func main() {
	f := missingRolls([]int{1, 2, 3, 4}, 6, 4)
	fmt.Println(f)
}

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	remainingSum := (m+n)*mean - sum
	if remainingSum > 6*n || remainingSum < n {
		return nil
	}
	res := make([]int, n)
	initValue := remainingSum / n
	remainingValue := remainingSum % n
	for i := range res {
		res[i] = initValue
		if remainingValue > 0 {
			res[i]++
			remainingValue--
		}
	}
	return res
}
