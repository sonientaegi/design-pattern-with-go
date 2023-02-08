package main

import "fmt"

type Abstraction[T any] struct {
	Impl Implementor[T]
}

func (s *Abstraction[T]) CallMethodA(value T) {
	s.Impl.MethodA(value)
}

func (s *Abstraction[T]) CallMethodB() {
	t := s.Impl.MethodB()
	fmt.Println("Value from Implementor is ", t)
}

type AbstractionSub0[T any] struct {
	Abstraction[T]
}

func (s *AbstractionSub0[T]) CallBothMethod(value T) {
	fmt.Println("Call both Method A and B at once")
	s.CallMethodA(value)
	s.CallMethodB()
}
