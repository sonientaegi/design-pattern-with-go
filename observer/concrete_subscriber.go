package main

import "fmt"

type ConcreteSubscriber struct {
	Name string
}

func (s *ConcreteSubscriber) Update(context any) {
	fmt.Println(s.Name, "has been notified.")
}
