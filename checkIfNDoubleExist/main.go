package main

func main() {
	println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}

func checkIfExist(arr []int) bool {
	exist := make(map[int]bool)
	for _, val := range arr {
		if exist[val*2] || (val%2 == 0 && exist[val/2]) {
			return true
		}
		exist[val] = true
	}
	return false
}
