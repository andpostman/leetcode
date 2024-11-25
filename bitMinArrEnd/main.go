package main

import "fmt"

func main() {
	fmt.Println(7 << 1)
	fmt.Println(minEnd(3, 4))
	fmt.Println(minEnd(2, 7))
}

func minEnd(n int, x int) int64 {
	// Store available bit positions that are 0 in x
	shifts := make([]int16, 0)

	// Initialize result with starting number x
	cur := int64(x)

	// Calculate how many 1's we need to add (n-1)
	calc := int64(n - 1)

	// Find all bit positions that are 0 in x (up to 32 bits)
	for i := 0; i < 32; i++ {
		if ((1 << i) & x) == 0 {
			shifts = append(shifts, int16(i))
		}
	}

	// Add remaining positions 32-63 since int64 is 64 bit
	for i := 32; i < 64; i++ {
		shifts = append(shifts, int16(i))
	}

	// Process each bit in calc (n-1)
	for i := 0; calc > 0; i++ {
		cur += (calc & 1) << shifts[i]
		calc >>= 1
	}

	return cur
}

// SECOND SOLUTION
func minEnd2(n int, x int) int64 {
	result := int64(x)
	remaining := int64(n - 1)
	position := int64(1)

	for remaining != 0 {
		if (int64(x) & position) == 0 {
			result |= (remaining & 1) * position
			remaining >>= 1
		}
		position <<= 1
	}

	return result
}
