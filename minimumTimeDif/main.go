package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"
)

func main() {
	fmt.Println(findMinDifference([]string{"00:00", "12:00", "23:59"}))
	fmt.Println(findMinDifference([]string{"02:45", "11:15", "18:30", "23:55"}))
}

func findMinDifference(timePoints []string) int {
	minutes := make([]int, len(timePoints))
	for i, point := range timePoints {
		if strings.Count(strings.Join(timePoints, " "), point) > 1 {
			return 0
		}
		from, err := time.Parse("15:04", point)
		if err != nil {
			log.Fatal(err)
		}
		minutes[i] = from.Minute() + from.Hour()*60
	}
	sort.Ints(minutes)
	min := math.MaxInt
	for i := 1; i < len(minutes); i++ {
		getMin(minutes[i], minutes[i-1], &min)
	}
	getMin(minutes[0], minutes[len(minutes)-1], &min)
	return min
}

func getMin(a, b int, min *int) {
	var diff int
	if a <= 12*60 || b <= 12*60 {
		res1 := 24*60 - int(math.Abs(float64(a-b)))
		res2 := int(math.Abs(float64(a - b)))
		if res1 < res2 {
			diff = res1
		} else {
			diff = res2
		}
	} else {
		diff = int(math.Abs(float64(a - b)))
	}

	if *min > diff {
		*min = diff
	}
}
