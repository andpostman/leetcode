package main

import (
	"fmt"
	"time"
)

func main() {
	fb := NewFizzBuzz(15)
	go fb.number(printNumber)
	go fb.fizz(printFizz)
	go fb.buzz(printBuzz)
	go fb.fizzbuzz(printFizzBuzz)
	defer close(fb.ch1)
	defer close(fb.ch2)
	defer close(fb.ch3)
	defer close(fb.ch4)
	time.Sleep(5 * time.Second)
}

type FizzBuzz struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
	ch4 chan struct{}
}

func NewFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{
		n:   n,
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}, 1),
		ch3: make(chan struct{}, 1),
		ch4: make(chan struct{}, 1),
	}
}

func (f *FizzBuzz) fizz(printF func()) {
	for range f.ch3 {
		printF()
		<-f.ch1
	}
}

func (f *FizzBuzz) buzz(printB func()) {
	for range f.ch2 {
		printB()
		<-f.ch1
	}
}

func (f *FizzBuzz) fizzbuzz(printFB func()) {
	for range f.ch4 {
		printFB()
		<-f.ch1
	}
}

func (f *FizzBuzz) number(printN func(int)) {
	for i := 1; i <= f.n; i++ {
		f.ch1 <- struct{}{}
		if i%3 == 0 && i%5 == 0 {
			f.ch4 <- struct{}{}
		} else if i%3 == 0 {
			f.ch3 <- struct{}{}
		} else if i%5 == 0 {
			f.ch2 <- struct{}{}
		} else {
			printN(i)
			<-f.ch1
		}
	}
}

func printFizz() {
	fmt.Print("\"fizz\"")
}

func printBuzz() {
	fmt.Print("\"buzz\"")
}

func printFizzBuzz() {
	fmt.Print("\"fizzbuzz\"")
}

func printNumber(x int) {
	fmt.Printf("\"%d\"", x)
}
