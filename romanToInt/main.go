package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func romanToInt(s string) int {
	romans := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := range s {
		sum += romans[s[i]]
		if i > 0 {
			if romans[s[i]] > romans[s[i-1]] {
				sum -= 2 * romans[s[i-1]]
			}
		}
	}
	return sum
}

// OR
func romanToInt2(s string) int {
	romans := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if result != 0 {
			switch v := romans[s[i]]; v {
			case 1:
				if romans[s[i+1]] == 5 || romans[s[i+1]] == 10 {
					result--
				} else {
					result++
				}
			case 10:
				if romans[s[i+1]] == 50 || romans[s[i+1]] == 100 {
					result -= 10
				} else {
					result += 10
				}
			case 100:
				if romans[s[i+1]] == 500 || romans[s[i+1]] == 1000 {
					result -= 100
				} else {
					result += 100
				}
			default:
				result += v
			}
		} else {
			result += romans[s[i]]
		}
	}
	return result
}
