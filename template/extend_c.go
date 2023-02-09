package main

import "fmt"

type ExtendC struct {
	Template
}

func NewExtendC() ExtendC {
	instance := ExtendC{
		Template: NewTemplate(),
	}
	instance.AbstractMethod = instance.methodC
	instance.HookMethod = instance.hookC

	return instance
}

func (s *ExtendC) methodC() {
	fmt.Println("This is implemented method of C.")
}

func (s *ExtendC) hookC() {
	fmt.Println("This is override hook of C.")
}
