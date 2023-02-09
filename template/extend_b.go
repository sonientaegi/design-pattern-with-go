package main

import "fmt"

type ExtendB struct {
	Template
}

func NewExtendB() ExtendB {
	instance := ExtendB{
		Template: NewTemplate(),
	}
	instance.AbstractMethod = instance.methodB

	return instance
}

func (s *ExtendB) methodB() {
	fmt.Println("This is implemented method of B.")
}
