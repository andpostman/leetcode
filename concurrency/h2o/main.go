package main

import (
	"fmt"
	"time"
)

func main() {
	h2o := NewH2O()
	go h2o.hydrogen(releaseHydrogen)
	go h2o.oxygen(releaseOxygen)
	go h2o.oxygen(releaseOxygen)
	go h2o.hydrogen(releaseHydrogen)
	go h2o.hydrogen(releaseHydrogen)
	go h2o.oxygen(releaseOxygen)
	go h2o.hydrogen(releaseHydrogen)
	go h2o.hydrogen(releaseHydrogen)
	go h2o.hydrogen(releaseHydrogen)
	time.Sleep(3 * time.Second)
}

type H2O struct {
	chH chan struct{}
	chO chan struct{}
}

func NewH2O() *H2O {
	return &H2O{chH: make(chan struct{}, 2), chO: make(chan struct{}, 2)}
}

func (h *H2O) hydrogen(print func()) {
	h.chH <- struct{}{}
	print()
	<-h.chO
}

func (h *H2O) oxygen(print func()) {
	h.chO <- struct{}{}
	h.chO <- struct{}{}
	print()
	<-h.chH
	<-h.chH
}

func releaseOxygen() {
	fmt.Print("O")
}

func releaseHydrogen() {
	fmt.Print("H")
}
