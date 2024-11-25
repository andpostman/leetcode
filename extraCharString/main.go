package main

import (
	"fmt"
)

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	fmt.Println(minExtraChar("dwmodizxvvbosxxw", []string{"ox", "lb", "diz", "gu", "v", "ksv", "o", "nuq", "r", "txhe", "e", "wmo", "cehy", "tskz", "ds", "kzbu"}))
}

func minExtraChar(s string, dictionary []string) int {
	dict := make(map[string]bool) // Use a map to store dictionary for O(1) lookup
	for _, word := range dictionary {
		dict[word] = true
	}

	n := len(s)
	dp := make([]int, n+1) // DP array initialized with maximum extra characters
	for i := 0; i <= n; i++ {
		dp[i] = n // Initialize DP array with maximum extra characters
	}
	dp[0] = 0 // No extra characters for an empty string

	// Iterate through the string
	for i := 1; i <= n; i++ {
		// Try all substrings ending at i
		for j := 0; j < i; j++ {
			sub := s[j:i] // Get substring s[j:i]
			if dict[sub] {
				dp[i] = min(dp[i], dp[j]) // If substring found in dictionary
			}
		}
		dp[i] = min(dp[i], dp[i-1]+1) // Consider current character as extra
	}

	return dp[n] // Return the result from dp[n]
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
