package main

import "fmt"

func main() {

	calendar := Constructor()

	fmt.Println(calendar.Book(10, 20))
	fmt.Println(calendar.Book(50, 60))
	fmt.Println(calendar.Book(10, 40))
	fmt.Println(calendar.Book(5, 15))
	fmt.Println(calendar.Book(5, 10))
	fmt.Println(calendar.Book(25, 55))

	//fmt.Println(calendar.Book(47, 50))
	//fmt.Println(calendar.Book(33, 41))
	//fmt.Println(calendar.Book(39, 45))
	//fmt.Println(calendar.Book(33, 42))
	//fmt.Println(calendar.Book(25, 32))
	//fmt.Println(calendar.Book(26, 35))
	//fmt.Println(calendar.Book(19, 25))
	//fmt.Println(calendar.Book(3, 8))
	//fmt.Println(calendar.Book(8, 13))
	//fmt.Println(calendar.Book(18, 27))
}

type MyCalendarTwo struct {
	booking       [][2]int
	doubleBooking [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		booking:       [][2]int{},
		doubleBooking: [][2]int{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	var doubleBooking [][2]int
	for _, ints := range this.booking {
		startBooking := ints[0]
		endBooking := ints[1]
		if start < endBooking && end > startBooking {

			for _, i := range this.doubleBooking {
				startDoubleBooking := i[0]
				endDoubleBooking := i[1]
				if start < endDoubleBooking && end > startDoubleBooking {
					return false
				}
			}
			startPosix, endPosix := 0, 0
			if start-startBooking > 0 {
				startPosix = start
			} else {
				startPosix = startBooking
			}
			if end-endBooking > 0 {
				endPosix = endBooking
			} else {
				endPosix = end
			}
			doubleBooking = append(doubleBooking, [2]int{startPosix, endPosix})
		}
	}
	this.booking = append(this.booking, [2]int{start, end})
	if len(doubleBooking) > 0 {
		this.doubleBooking = append(this.doubleBooking, doubleBooking...)
	}
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
