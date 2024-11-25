package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getLucky("zbax", 2))
}

func getLucky(s string, k int) int {
	var alphabetArr []string
	for _, i := range s {
		number := fmt.Sprint(i - 'a' + 1)
		numbers := strings.Split(number, "")
		alphabetArr = append(alphabetArr, numbers...)
	}
	for i := 0; i < k; i++ {
		sum := 0
		for _, element := range alphabetArr {
			el, _ := strconv.Atoi(element)
			sum += el
		}
		number := fmt.Sprint(sum)
		alphabetArr = strings.Split(number, "")
	}
	resNumber := strings.Join(alphabetArr, "")
	res, _ := strconv.Atoi(resNumber)
	return res
}
