package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maximumSwap(1993))
}

func maximumSwap(num int) int {
	s := fmt.Sprint(num)
	builder := strings.Builder{}
	builder.WriteString(s)
	max := s
	for i := 0; i < len(s); i++ {

		charI := s[i]
		maxElement := charI
		maxIndex := i

		for j := len(s) - 1; j > i; j-- {
			charJ := s[j]
			if charI < charJ && maxElement < charJ {
				maxElement = charJ
				maxIndex = j
			}
		}
		num := swap(s, i, maxIndex)
		if strings.Compare(num, max) > 0 {
			max = num
		}
	}
	res, _ := strconv.Atoi(max)
	return res
}

func swap(s string, i1, i2 int) string {
	out := []rune(s)
	out[i1], out[i2] = out[i2], out[i1]
	return string(out)
}
