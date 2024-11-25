package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	start := 0
	end := len(s) - 1
	var processedStart, processedEnd []string
	bestLenStart, bestLenEnd := 0, 0
	expression := true
	startExpression, endExpression := true, true
	nextStart, nextEnd := 0, 0
	for i, k := start, end; expression; {
		expression = i <= k
		if !expression {
			if startExpression {
				startExpression = len(processedStart) != 0 && len(s) > i && 0 <= k
			}
			if endExpression {
				endExpression = len(processedEnd) != 0 && len(s) > i && 0 <= k
			}
			if expression = startExpression || endExpression; !expression {
				break
			}
		}
		charStart := fmt.Sprintf("%c", s[i])
		charEnd := fmt.Sprintf("%c", s[k])
		if len(processedStart) == 1 {
			nextStart = i
		}
		if len(processedEnd) == 1 {
			nextEnd = k
		}

		processedStart = checkArrToAdd(charStart, processedStart, &bestLenStart)
		processedEnd = checkArrToAdd(charEnd, processedEnd, &bestLenEnd)

		if len(processedStart) != 0 {
			i++
		} else {
			i = nextStart
		}

		if len(processedEnd) != 0 {
			k--
		} else {
			k = nextEnd
		}
	}
	if bestLenStart < bestLenEnd {
		return bestLenEnd
	} else {
		return bestLenStart
	}
}

func checkArrToAdd(char string, arr []string, bestLen *int) []string {
	if arr == nil || !strings.Contains(strings.Join(arr, " "), char) {
		arr = append(arr, char)
	} else {
		arr = []string{}
	}
	if *bestLen < len(arr) {
		*bestLen = len(arr)
	}
	return arr
}
