package main

import "fmt"

func main() {
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}

func decrypt(code []int, k int) []int {
	decryptedCode := make([]int, len(code))
	for i := range decryptedCode {
		if k < 0 {
			for j, r := i-1, k; r < 0; r, j = r+1, j-1 {
				if j < 0 {
					j = len(code) - 1
				}
				decryptedCode[i] += code[j]
			}
		} else if k > 0 {
			for j, r := i+1, k; r > 0; r, j = r-1, j+1 {
				if j >= len(code) {
					j = 0
				}
				decryptedCode[i] += code[j]
			}
		}
	}
	return decryptedCode
}
