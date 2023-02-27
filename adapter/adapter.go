package main

import (
	"fmt"
)

type Target struct{}

func (s *Target) Call() {
	fmt.Println("Target - call")
}

type Adaptee struct {
	isInvoked bool
}

func (s *Adaptee) Invoke() {
	fmt.Println("Adaptee - invoke")
	s.isInvoked = true
}

func (s *Adaptee) Process() {
	if s.isInvoked {
		fmt.Println("Adaptee - execution")
	}
}

type Adapter struct {
	adaptee Adaptee
}

func (s *Adapter) Call() {
	fmt.Println("Adapter - call")
	s.adaptee.Invoke()
	s.adaptee.Process()
	fmt.Println("Adapter - done")
}
