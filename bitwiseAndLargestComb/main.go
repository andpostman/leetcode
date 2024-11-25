package main

func main() {
	largestCombination([]int{16, 17, 71, 62, 12, 24, 14})
	largestCombination([]int{1024, 1, 512, 2097152, 8, 32, 256, 4096, 16, 4194304, 524288, 8388608, 2048, 64, 8192, 4, 16384, 262144, 2, 131072, 1048576, 128, 32768, 65536})
}

func largestCombination(candidates []int) int {
	m := make([]int, 24)
	lenMax := 0
	for _, candidate := range candidates {
		lenMax = max(lenMax, findBitPosition(uint(candidate), m))
	}
	return lenMax
}

func findBitPosition(num uint, m []int) int {
	lenMax := 0
	for i := range m {
		if (num>>i)&1 == 1 {
			m[i]++
			lenMax = max(lenMax, m[i])
		}
	}
	return lenMax
}
