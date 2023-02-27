package main

import "fmt"

type Encapsulate interface {
	Method()
}

type AbstractObject struct {
	Capsule Encapsulate
}

func (s *AbstractObject) Action() {
	if s.Capsule == nil {
		fmt.Println("Do nothing")
	} else {
		s.Capsule.Method()
	}
}
