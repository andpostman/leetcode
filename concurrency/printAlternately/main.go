package main

import (
	"fmt"
	"time"
)

func main() {
	fooBar := NewFooBar(5)
	go fooBar.Foo(printFoo)
	go fooBar.Bar(printBar)
	time.Sleep(1 * time.Second)
}

type FooBar struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:   n,
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}, 1),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.ch1 <- struct{}{}
		printFoo()
		fb.ch2 <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch2
		printBar()
		<-fb.ch1
	}
}

func printFoo() {
	fmt.Print("foo")
}

func printBar() {
	fmt.Print("bar")
}
