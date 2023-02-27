package main

import "fmt"

type Component interface {
	Operate()
}

type ConcreteComponent struct {
	// Component
	Description string
}

func (s *ConcreteComponent) Operate() {
	fmt.Print("Operate IN > ")
	fmt.Print(s.Description)
	fmt.Print(" > Operate OUT > ")
}
