package main

import "fmt"

type ExtendA struct {
	Template
}

func NewExtendA() ExtendA {
	instance := ExtendA{
		Template: NewTemplate(),
	}
	instance.AbstractMethod = instance.methodA

	return instance
}

func (s *ExtendA) methodA() {
	fmt.Println("This is implemented method of A.")
}
