package main

func main() {
	println(takeCharacters("aabaaaacaabc", 2))
	println(takeCharacters("abc", 1))
	println(takeCharacters("cbbac", 1))
}

func takeCharacters(s string, k int) int {
	count := [3]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	for i := range count {
		if count[i] < k {
			return -1
		}
	}
	window := [3]int{}
	l, maxWindow := 0, 0
	for r := range s {
		window[s[r]-'a']++
		for l <= r &&
			(count[0]-window[0] < k ||
				count[1]-window[1] < k ||
				count[2]-window[2] < k) {
			window[s[l]-'a']--
			l++
		}
		maxWindow = max(maxWindow, r-l+1)
	}
	return len(s) - maxWindow
}
