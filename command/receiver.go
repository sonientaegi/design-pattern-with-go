package main

import "fmt"

type Receiver struct {
	Name string
}

func (s *Receiver) DoSomething() {
	fmt.Println("Hello! I am", s.Name)
}
