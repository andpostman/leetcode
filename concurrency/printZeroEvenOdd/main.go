package main

import (
	"fmt"
	"time"
)

func main() {
	zeo := NewZeroEvenOdd(10)
	go zeo.zero(printNumber)
	go zeo.even(printNumber)
	go zeo.odd(printNumber)
	time.Sleep(1 * time.Second)
}

type ZeroEvenOdd struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
	ch3 chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	return &ZeroEvenOdd{
		n:   n,
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}, 1),
		ch3: make(chan struct{}, 1),
	}
}

func (z *ZeroEvenOdd) zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		z.ch1 <- struct{}{}
		printNumber(0)
		if i%2 == 0 {
			z.ch2 <- struct{}{}
		} else {
			z.ch3 <- struct{}{}
		}

	}
}

func (z *ZeroEvenOdd) odd(printNumber func(int)) {
	for i := 1; i <= z.n; i += 2 {
		<-z.ch3
		printNumber(i)
		<-z.ch1
	}
}

func (z *ZeroEvenOdd) even(printNumber func(int)) {
	for i := 2; i <= z.n; i += 2 {
		<-z.ch2
		printNumber(i)
		<-z.ch1
	}
}

func printNumber(n int) {
	fmt.Print(n)
}
