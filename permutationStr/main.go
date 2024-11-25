package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count, s2Count := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if s1Count == s2Count {
			return true
		}
		s2Count[s2[i]-'a']--
		s2Count[s2[i+len(s1)]-'a']++
	}
	return s1Count == s2Count
}

// SECOND sliding window
func checkInclusion2(s1 string, s2 string) bool {
	l, count := 0, [26]int{}
	for i := range s1 {
		count[s1[i]-97]++
	}

	for r := range s2 {
		count[s2[r]-97]--
		if count == [26]int{} {
			return true
		}

		if r+1 >= len(s1) {
			count[s2[l]-97]++
			l++
		}
	}
	return false
}
