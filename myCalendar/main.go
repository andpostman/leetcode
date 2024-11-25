package main

import "fmt"

func main() {

	calendar := Constructor()
	fmt.Println(calendar.Book(47, 50))
	fmt.Println(calendar.Book(33, 41))
	fmt.Println(calendar.Book(39, 45))
	fmt.Println(calendar.Book(33, 42))
	fmt.Println(calendar.Book(25, 32))
	fmt.Println(calendar.Book(26, 35))
	fmt.Println(calendar.Book(19, 25))
	fmt.Println(calendar.Book(3, 8))
	fmt.Println(calendar.Book(8, 13))
	fmt.Println(calendar.Book(18, 27))
}

type MyCalendar struct {
	booked [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{booked: [][2]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, ints := range this.booked {
		if start < ints[1] && end > ints[0] {
			return false
		}
	}
	this.booked = append(this.booked, [2]int{start, end})
	return true
}
