package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	foo := &Foo{}
	go foo.first(ch)
	go foo.second(ch)
	go foo.third(ch)
	time.Sleep(3 * time.Second)
}

type Foo struct{}

func (f *Foo) first(ch chan struct{}) {
	fmt.Print("first")
	ch <- struct{}{}
}

func (f *Foo) second(ch chan struct{}) {
	<-ch
	fmt.Print("second")
	ch <- struct{}{}
}

func (f *Foo) third(ch chan struct{}) {
	time.Sleep(1 * time.Second)
	<-ch
	fmt.Print("third")
}
